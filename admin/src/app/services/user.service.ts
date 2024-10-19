import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';

@Injectable({
  providedIn: 'root',
})
export class UserService {
  BASE_URL = 'http://localhost:5051/api/v1/iam'; // TODO: change the base url to api gateway instead of direct server
  apiToken = '';

  constructor(private http: HttpClient) {}

  login(username: string, password: string) {
    return this.http.post(`${this.BASE_URL}/login`, { username, password });
  }

  setApiToken(apiToken: string) {
    this.apiToken = apiToken;
  }

  isLogin() {
    return !!this.apiToken;
  }
}
