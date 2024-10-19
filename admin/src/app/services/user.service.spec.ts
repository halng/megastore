import { TestBed } from '@angular/core/testing';
import { UserService } from './user.service';
import {
  HttpTestingController,
} from '@angular/common/http/testing';

import { provideHttpClient } from '@angular/common/http';
import { provideHttpClientTesting } from '@angular/common/http/testing';

describe('UserService', () => {
  let service: UserService;
  let httpMock: HttpTestingController;

  beforeEach(() => {
    TestBed.configureTestingModule({
      imports: [],
      providers: [UserService, provideHttpClient(), provideHttpClientTesting()],
    });
    service = TestBed.inject(UserService);
    httpMock = TestBed.inject(HttpTestingController);
  });

  afterEach(() => {
    httpMock.verify();
  });

  it('should be created', () => {
    expect(service).toBeTruthy();
  });

  it('should login successfully', () => {
    const mockResponse = { token: '12345' };
    const username = 'testuser';
    const password = 'testpass';

    service.login(username, password).subscribe((response: any) => {
      expect(response.token).toEqual(mockResponse.token);
    });

    const req = httpMock.expectOne(`${service.BASE_URL}/login`);
    expect(req.request.method).toBe('POST');
    expect(req.request.body).toEqual({ username, password });
    req.flush(mockResponse);
  });

  it('should handle 404 error on login', () => {
    const username = 'testuser';
    const password = 'testpass';

    service.login(username, password).subscribe(
      () => fail('should have failed with 404 error'),
      (error) => {
        expect(error.status).toBe(404);
      }
    );

    const req = httpMock.expectOne(`${service.BASE_URL}/login`);
    expect(req.request.method).toBe('POST');
    req.flush('Login failed', { status: 404, statusText: 'Not Found' });
  });

  it('should handle 401 error on login', () => {
    const username = 'testuser';
    const password = 'testpass';

    service.login(username, password).subscribe(
      () => fail('should have failed with 401 error'),
      (error) => {
        expect(error.status).toBe(401);
      }
    );

    const req = httpMock.expectOne(`${service.BASE_URL}/login`);
    expect(req.request.method).toBe('POST');
    req.flush('Unauthorized', { status: 401, statusText: 'Unauthorized' });
  });

  it('should set API token', () => {
    const token = '12345';
    service.setApiToken(token);
    expect(service.apiToken).toBe(token);
  });

  it('should return true if logged in', () => {
    service.setApiToken('12345');
    expect(service.isLogin()).toBeTrue();
  });

  it('should return false if not logged in', () => {
    service.setApiToken('');
    expect(service.isLogin()).toBeFalse();
  });
});
