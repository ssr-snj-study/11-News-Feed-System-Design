from dependency_injector import containers, providers
from api.v1.auth.repository import AuthRepository
from api.v1.auth.services import AuthServices


class Container(containers.DeclarativeContainer):
    # logger
    logger = providers.Singleton()

    # PostgreSQL 리소스
    postgres_engine = providers.Resource()

    # Redis 리소스
    redis_client = providers.Resource()

    # RabbitMQ 리소스
    rabbimq_connection = providers.Factory()

    # Auth repository
    auth_repository = providers.Factory(
        AuthRepository, logger=logger, rdb_session=postgres_engine.provided.get_pg_session
    )

    # Auth services
    auth_services = providers.Factory(
        AuthServices, logger=logger, auth_repository=auth_repository, rabbitmq=rabbimq_connection
    )
