from fastapi import APIRouter, Depends
from dependency_injector.wiring import inject, Provide
from .services import FeedServices

router = APIRouter(prefix="/feed", tags=["Feed"])

@router.get("/")
@inject
async def get_feeds(
    feed_services: FeedServices = Depends(Provide["feed_container.feed_services"])
):
    feeds = await feed_services.get_feeds()
    return {"feeds": feeds} 

@router.post("/")
@inject
async def create_feed(
    title: str,
    content: str,
    author: str,
    feed_services: FeedServices = Depends(Provide["feed_container.feed_services"])
):
    feed_id = await feed_services.create_feed(title, content, author)
    return {"message": "Feed created successfully", "feed_id": feed_id}