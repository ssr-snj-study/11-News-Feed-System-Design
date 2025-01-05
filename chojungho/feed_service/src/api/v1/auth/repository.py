from sqlalchemy.ext.asyncio import AsyncSession
from sqlalchemy.future import select
from infrastructure.schema.user_schema import UserInfo

class AuthRepository:
    def __init__(self, rdb_session: AsyncSession):
        self._rdb_session = rdb_session

    async def save_tokens(self, user_id: str, access_token: str, refresh_token: str):
        async with self._rdb_session() as session:
            user_info = await session.execute(select(UserInfo).where(UserInfo.user_id == user_id))
            user = user_info.scalar_one_or_none()
            if user:
                user.access_token = access_token
                user.refresh_token = refresh_token
            else:
                user = UserInfo(user_id=user_id, access_token=access_token, refresh_token=refresh_token)
                session.add(user)
            await session.commit()
