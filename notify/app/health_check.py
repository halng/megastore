from fastapi import status
from pydantic import BaseModel


class HealthCheckResponse(BaseModel):
    success: str = "UP"


