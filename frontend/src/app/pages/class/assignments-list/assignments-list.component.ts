import { Component, Injectable, OnInit } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { ActivatedRoute } from '@angular/router';
import { Observable } from 'rxjs';
import { RoleService } from 'src/app/services/role.service';

interface Assignment {
  id: number;
  name: string;
  mark: number;
}

@Injectable({
  providedIn: 'root',
})
export class AssignmentsService {
  private baseUrl = 'http://localhost:8080/classes/assignments-list';
  private markUrl = 'http://localhost:8080/classes/mark';

  constructor(private http: HttpClient) {}

  getAssignments(classId: string): Observable<Assignment[]> {
    return this.http.get<Assignment[]>(`${this.baseUrl}/${classId}`, { withCredentials: true });
  }

  getGrade(assignmentId: number): Observable<{ mark: number }> {
    return this.http.get<{ mark: number }>(
      `${this.markUrl}/${assignmentId}`,
      { withCredentials: true }
    );
  }
}

@Component({
  selector: 'app-assignments-list',
  templateUrl: './assignments-list.component.html',
  styleUrls: ['./assignments-list.component.css']
})
export class AssignmentsListComponent implements OnInit {
  assignments: Assignment[] = [];
  isModalOpen = false;
  isTeacher = false;

  constructor(
    private route: ActivatedRoute,
    private assignmentsService: AssignmentsService,
    private roleService: RoleService
  ) {}

  ngOnInit(): void {
    const classID = this.route.snapshot.paramMap.get('id');
    console.log(classID)
    if (classID) {
      this.assignmentsService.getAssignments(classID).subscribe((data) => {
        this.assignments = data;

        this.roleService.isTeacher(Number(classID)).subscribe(res => {
          this.isTeacher = res.isTeacher;
        });
        
        this.assignments.forEach((assignment) => {
          assignment.mark = 0;
          this.assignmentsService.getGrade(assignment.id).subscribe((gradeResponse) => {
            assignment.mark = gradeResponse.mark;
            console.log(gradeResponse);
          });
        });
      });
    }
  }

  openModal(): void {
    this.isModalOpen = true;
  }

  closeModal(): void {
    this.isModalOpen = false;
  }
}
