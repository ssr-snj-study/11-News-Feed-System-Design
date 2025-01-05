from dependency_injector import containers, providers
from .repository import FeedRepository
from .services import FeedServices

class FeedContainer(containers.DeclarativeContainer):
    config = providers.Configuration()

    # PostgreSQL 리소스
    postgres_engine = providers.Resource()

    # Redis 리소스
    redis_client = providers.Resource()

    # Feed repository
    feed_repository = providers.Factory(
        FeedRepository, rdb_session=postgres_engine.provided.get_pg_session
    )

    # Feed services
    feed_services = providers.Factory(
        FeedServices, feed_repository=feed_repository, redis_client=redis_client
    )
