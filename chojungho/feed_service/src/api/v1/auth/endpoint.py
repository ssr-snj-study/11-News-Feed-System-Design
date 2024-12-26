from fastapi import APIRouter, Request
from dependency_injector.wiring import inject, Provide
from fastapi.templating import Jinja2Templates
import urllib.parse
from starlette.responses import RedirectResponse
from fastapi.responses import JSONResponse
import requests

router = APIRouter(prefix="/auth", tags=["Auth"])
templates = Jinja2Templates(directory="templates/auth")


@router.get("/login")
async def get_login_view(request: Request):
    return templates.TemplateResponse(request=request, name="login.html")


@router.get("/google_login")
@inject
async def google_login(config=Provide["config"]) -> RedirectResponse:
    # Google OAuth 2.0 인증 URL 생성
    base_url = "https://accounts.google.com/o/oauth2/auth"
    params = {
        "client_id": config["GOOGLE_CLIENT_ID"],
        "redirect_uri": config["GOOGLE_REDIRECT_URI"],
        "response_type": "code",
        "scope": "email profile",
        "access_type": "offline",
    }
    login_url = f"{base_url}?{urllib.parse.urlencode(params)}"
    return RedirectResponse(url=login_url, status_code=302)


@router.get("/google_callback")
@inject
async def google_callback_view(request: Request, config=Provide["config"]):
    code = request.query_params.get("code")
    if not code:
        return JSONResponse(content={"error": "No code provided"}, status_code=400)

    # 구글에 액세스 토큰 요청
    token_url = "https://oauth2.googleapis.com/token"
    token_data = {
        "code": code,
        "client_id": config["GOOGLE_CLIENT_ID"],
        "client_secret": config["GOOGLE_CLIENT_SECRET"],
        "redirect_uri": config["GOOGLE_REDIRECT_URI"],
        "grant_type": "authorization_code",
    }
    token_response = requests.post(token_url, data=token_data)
    token_json = token_response.json()
    access_token = token_json.get("access_token")

    if not access_token:
        return JSONResponse(content={"error": "Failed to obtain access token"}, status_code=400)

    # 구글 API로 사용자 정보 요청
    # todo: access_token 을 DB에 추가 및 refresh_token 기능도 추가 필요함
    user_info_url = "https://www.googleapis.com/oauth2/v1/userinfo"
    user_info_response = requests.get(user_info_url, headers={"Authorization": f"Bearer {access_token}"})
    user_info = user_info_response.json()

    # 사용자 정보 표시
    return templates.TemplateResponse(request=request, name="login_user.html", context={"user_info": user_info})
