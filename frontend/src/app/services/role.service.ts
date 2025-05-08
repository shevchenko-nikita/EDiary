import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { Observable, map } from 'rxjs';

@Injectable({
  providedIn: 'root'
})
export class RoleService {
  constructor(private http: HttpClient) {}

  isTeacher(classID: number) {
    return this.http.get<{ isTeacher: boolean }>(
      `http://localhost:8080/classes/is-teacher/${classID}`, 
      { withCredentials: true }
    );
  }
}
