import 'dart:convert';

import 'package:flutter/material.dart';
import 'package:firebase_core/firebase_core.dart';
import 'package:firebase_messaging/firebase_messaging.dart';
import 'notification_service.dart';  // 로컬 알림 서비스 import
import 'notification_handler.dart'; // 메시지 핸들링 서비스 import
import 'firebase_options.dart';  // Firebase 옵션 import
import 'package:dio/dio.dart';
import 'package:cookie_jar/cookie_jar.dart';
import 'package:dio_cookie_manager/dio_cookie_manager.dart';

final Dio dio = Dio(BaseOptions(baseUrl: 'http://10.0.2.2:1323/',headers: {"Content-Type": "application/json"}));
final CookieJar cookieJar = CookieJar();
const String userId = "test";

void main() async {
  WidgetsFlutterBinding.ensureInitialized();
  // Firebase 초기화
  await Firebase.initializeApp(options: DefaultFirebaseOptions.currentPlatform);

  // Firebase Messaging 인스턴스 생성
  FirebaseMessaging messaging = FirebaseMessaging.instance;
  NotificationSettings settings = await messaging.requestPermission(
    alert: true,
    badge: true,
    sound: true,
  );

  // FCM 토큰 확인 및 로그로 출력
  String? token = await messaging.getToken();
  if (token != null) {
    print("FCM 토큰: $token");
    sendTokenToServer(token);
  } else {
    print("FCM 토큰을 가져올 수 없습니다.");
  }

  // Firebase 메시징 초기화
  FirebaseMessaging.onBackgroundMessage(NotificationHandler.backgroundHandler);
  FirebaseMessaging.onMessage.listen(NotificationHandler.foregroundHandler);

  setupDio();

  await login();

  runApp(MyApp());
}

class MyApp extends StatelessWidget {

  @override
  Widget build(BuildContext context) {
    return MaterialApp(
      home: PostRequestScreen(),
    );
  }
}

class PostRequestScreen extends StatefulWidget {
  @override
  _PostRequestScreenState createState() => _PostRequestScreenState();
}

class _PostRequestScreenState extends State<PostRequestScreen> {
  // TextEditingController로 입력 필드 값 관리
  final TextEditingController _nameController = TextEditingController();
  final TextEditingController _feedController = TextEditingController();

  // POST 요청 함수
  Future<void> sendPostRequest(String name, String feed) async {
    try {
      final response = await dio.post('/api/v1/feed', data: {"name":name, "contents": feed});
      print('요청 헤더: ${response.requestOptions.headers}');

      if (response.statusCode == 200 || response.statusCode == 201) {
        // 요청 성공
        print('POST 요청 성공: ${response.data}');
        showDialog(
          context: context,
          builder: (context) => AlertDialog(
            title: Text('성공'),
            content: Text('데이터가 성공적으로 전송되었습니다.'),
            actions: [
              TextButton(
                onPressed: () => Navigator.pop(context),
                child: Text('확인'),
              ),
            ],
          ),
        );
      } else {
        // 요청 실패
        print('POST 요청 실패: ${response.statusCode}');
        showDialog(
          context: context,
          builder: (context) => AlertDialog(
            title: Text('실패'),
            content: Text('서버 오류가 발생했습니다: ${response.statusCode}'),
            actions: [
              TextButton(
                onPressed: () => Navigator.pop(context),
                child: Text('확인'),
              ),
            ],
          ),
        );
      }
    } catch (error) {
      // 네트워크 오류 처리
      print('POST 요청 중 오류 발생: $error');
      showDialog(
        context: context,
        builder: (context) => AlertDialog(
          title: Text('오류'),
          content: Text('네트워크 오류가 발생했습니다.'),
          actions: [
            TextButton(
              onPressed: () => Navigator.pop(context),
              child: Text('확인'),
            ),
          ],
        ),
      );
    }
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        title: Text('HTTP POST 예제'),
      ),
      body: Padding(
        padding: const EdgeInsets.all(16.0),
        child: Column(
          children: [
            TextField(
              controller: _nameController,
              decoration: InputDecoration(labelText: '이름'),
            ),
            TextField(
              controller: _feedController,
              decoration: InputDecoration(labelText: '내용'),
            ),
            SizedBox(height: 20),
            ElevatedButton(
              onPressed: () {
                final name = _nameController.text;
                final feed = _feedController.text;
                sendPostRequest(name, feed); // POST 요청 호출
              },
              child: Text('posting'),
            ),
          ],
        ),
      ),
    );
  }
}

Future<void> login() async {
  try {
    final response = await dio.post('/auth', data: {
      'name': userId,
    });

    print('로그인 응답: ${response.data}');
    print('Set-Cookie 헤더로 받은 쿠키: ${response.headers['set-cookie']}');

    // 쿠키 전송 확인
    final cookies = await cookieJar.loadForRequest(Uri.parse('http://10.0.2.2:1323/'));
    print('요청 시 쿠키: $cookies');
  } catch (e) {
    print('로그인 중 오류 발생: $e');
  }
}

Future<void> sendTokenToServer(String token) async {
  try {
    final response = await dio.post(
      "/SetDevice", // 경로 수정
      data: {"token": token, "name": userId},
    );

    if (response.statusCode == 200) {
      print("Token sent successfully");
    } else {
      print("Failed to send token: ${response.statusCode}");
    }
  } catch (e) {
    print('토큰 전송 중 오류 발생: $e');
  }
}

// 요청 시 쿠키를 자동으로 관리하려면 쿠키 관리자를 추가합니다.
void setupDio() {
  dio.interceptors.add(CookieManager(cookieJar)); // 쿠키 관리자를 인터셉터로 추가
}