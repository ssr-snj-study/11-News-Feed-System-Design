from contextlib import AbstractAsyncContextManager
from typing import Callable
from sqlalchemy.ext.asyncio import AsyncSession
from sqlalchemy.future import select
from infrastructure.schema.feed_schema import Feed

class FeedRepository:
    def __init__(self, rdb_session: Callable[..., AbstractAsyncContextManager[AsyncSession]]):
        self._rdb_session = rdb_session

    async def save_feed(self, title: str, content: str, author: str):
        async with self._rdb_session() as session:
            new_feed = Feed(title=title, content=content, author=author)
            session.add(new_feed)
            await session.commit()
            return new_feed.id 

    async def get_feeds(self):
        async with self._rdb_session() as session:
            result = await session.execute(select(Feed))
            return result.scalars().all() 