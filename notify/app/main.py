from fastapi import FastAPI, status

from app.health_check import HealthCheckResponse

app = FastAPI()


@app.get("/health", tags=["HealthCheck"])
def get_health():
    return HealthCheckResponse(status="OK")
