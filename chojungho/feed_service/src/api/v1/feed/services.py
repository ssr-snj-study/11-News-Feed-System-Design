from redis.asyncio import Redis
from .repository import FeedRepository
from dependency_injector.wiring import inject
import json

class FeedServices:
    @inject
    def __init__(self, feed_repository: FeedRepository, redis_client: Redis):
        self._feed_repository = feed_repository
        self._redis_client = redis_client
    
    async def create_feed(self, title: str, content: str, author: str):
        feed_id = await self._feed_repository.save_feed(title, content, author)
        feed_data = {"title": title, "content": content, "author": author}
        await self._redis_client.set(f"feed:{feed_id}", json.dumps(feed_data))
        return feed_id 

    async def get_feeds(self):
        # Redis 캐시에서 피드 조회
        cached_feeds = await self._redis_client.get("feeds")
        if cached_feeds:
            return json.loads(cached_feeds)

        # RDB에서 피드 조회
        feeds = await self._feed_repository.get_feeds()
        feeds_data = [{"id": feed.id, "title": feed.title, "content": feed.content, "author": feed.author} for feed in feeds]

        # Redis에 캐시 저장
        await self._redis_client.set("feeds", json.dumps(feeds_data))
        return feeds_data 