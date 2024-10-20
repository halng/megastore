import { TestBed } from '@angular/core/testing';
import { UserService } from './user.service';
import {
  HttpTestingController,
} from '@angular/common/http/testing';

import { provideHttpClient } from '@angular/common/http';
import { provideHttpClientTesting } from '@angular/common/http/testing';
import { UserCreate } from '../types';

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
  
  it('should create staff successfully', () => {
    const mockData: UserCreate = { username: 'testuser', email: 'test@example.com', firstName: 'Test', lastName: 'User' };
    const mockResponse = { success: true };

    localStorage.setItem(service.authKey, JSON.stringify({ 'api-token': 'test-token', 'id': 'test-id' }));

    service.createStaff(mockData).subscribe(response => {
      expect(response).toEqual(mockResponse);
    });

    const req = httpMock.expectOne(`${service.BASE_URL}/create-staff`);
    expect(req.request.method).toBe('POST');
    expect(req.request.headers.get('X-API-SECRET-TOKEN')).toBe('test-token');
    expect(req.request.headers.get('X-API-USER-ID')).toBe('test-id');
    req.flush(mockResponse);
  });

  it('should throw error if no auth token found', () => {
    localStorage.removeItem(service.authKey);

    expect(() => service.createStaff({ username: 'testuser', email: 'test@example.com', firstName: 'Test', lastName: 'User' }))
      .toThrow(new Error('No auth token found'));
  });

  it('should remove auth keys on logout', () => {
    spyOn(localStorage, 'removeItem');

    service.logout();

    expect(localStorage.removeItem).toHaveBeenCalledWith(service.authKey);
    expect(localStorage.removeItem).toHaveBeenCalledWith(service.authExpireKey);
  });

  
});
