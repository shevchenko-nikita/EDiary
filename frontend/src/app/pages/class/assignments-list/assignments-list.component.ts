import { Component, Injectable, OnInit } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { ActivatedRoute } from '@angular/router';
import { Observable } from 'rxjs';
import { RoleService } from 'src/app/services/role.service';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';

interface Assignment {
  id: number;
  name: string;
  mark: number;
  deadLine?: string;
}

interface CreateAssignmentRequest {
  name: string;
  class_id: number;
  time_created: string; // формат "yyyy-mm-dd hh:mm:ss"
  deadline: string;   // формат "yyyy-mm-dd hh:mm:ss"
}

@Injectable({
  providedIn: 'root',
})
export class AssignmentsService {
  private baseUrl = 'http://localhost:8080/classes/assignments-list';
  private markUrl = 'http://localhost:8080/classes/mark';
  private createUrl = 'http://localhost:8080/classes/create-assignment';

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

  createAssignment(request: CreateAssignmentRequest): Observable<any> {
    return this.http.post(this.createUrl, request, { withCredentials: true });
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
  assignmentForm: FormGroup;
  currentClassId: number = 0;
  
  // Получаем текущую дату и время
  currentDate = new Date();
  
  // Устанавливаем минимальную дату для выбора дедлайна (текущая дата)
  minDeadlineDate = this.formatDateForInput(this.currentDate);
  
  // Устанавливаем минимальное время для выбора (текущее время)
  minDeadlineTime = this.formatTimeForInput(this.currentDate);

  constructor(
    private route: ActivatedRoute,
    private assignmentsService: AssignmentsService,
    private roleService: RoleService,
    private fb: FormBuilder
  ) {
    // Получаем дату через неделю для дедлайна по умолчанию
    const nextWeek = new Date();
    nextWeek.setDate(nextWeek.getDate() + 7);
    
    this.assignmentForm = this.fb.group({
      taskName: ['', Validators.required],
      deadlineDate: [this.formatDateForInput(nextWeek), Validators.required],
      deadlineTime: ['23:59', Validators.required]
    });
  }

  ngOnInit(): void {
    const classID = this.route.snapshot.paramMap.get('id');
    console.log(classID);
    if (classID) {
      this.currentClassId = Number(classID);
      this.assignmentsService.getAssignments(classID).subscribe((data) => {
        this.assignments = data.reverse();

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
    this.assignmentForm.reset({
      taskName: '',
      deadlineDate: this.formatDateForInput(this.getNextWeekDate()),
      deadlineTime: '23:59'
    });
  }
  
  // Вспомогательные методы для работы с датами
  formatDateForInput(date: Date): string {
    return date.toISOString().split('T')[0];
  }
  
  formatTimeForInput(date: Date): string {
    return date.toTimeString().substring(0, 5);
  }
  
  getNextWeekDate(): Date {
    const nextWeek = new Date();
    nextWeek.setDate(nextWeek.getDate() + 7);
    return nextWeek;
  }

  // Преобразование выбранной даты и времени в формат "yyyy-mm-dd hh:mm:ss" для отправки на сервер
  combineDateAndTime(date: string, time: string): string {
    return `${date} ${time}:00`;
  }

  // Форматирование текущей даты и времени в формат "yyyy-mm-dd hh:mm:ss"
  formatDateTimeForServer(date: Date): string {
    const year = date.getFullYear();
    const month = String(date.getMonth() + 1).padStart(2, '0');
    const day = String(date.getDate()).padStart(2, '0');
    const hours = String(date.getHours()).padStart(2, '0');
    const minutes = String(date.getMinutes()).padStart(2, '0');
    
    return `${year}-${month}-${day} ${hours}:${minutes}:00`;
  }

  addAssignment(): void {
    if (this.assignmentForm.valid) {
      // Текущая дата и время для time_created в формате "yyyy-mm-dd hh:mm:ss"
      const currentDateTime = this.formatDateTimeForServer(new Date());
      
      // Комбинируем дату и время дедлайна в формате "yyyy-mm-dd hh:mm:ss"
      const deadlineDateTime = this.combineDateAndTime(
        this.assignmentForm.value.deadlineDate,
        this.assignmentForm.value.deadlineTime
      );

      const request: CreateAssignmentRequest = {
        name: this.assignmentForm.value.taskName,
        class_id: this.currentClassId,
        time_created: currentDateTime,
        deadline: deadlineDateTime
      };

      this.assignmentsService.createAssignment(request).subscribe({
        next: (response) => {
          console.log('Задание успешно создано', response);
          this.closeModal();
          // Обновляем список заданий
          this.assignmentsService.getAssignments(String(this.currentClassId)).subscribe((data) => {
            this.assignments = data.reverse();
            
            // Обновляем оценки для новых заданий
            this.assignments.forEach((assignment) => {
              assignment.mark = 0;
              this.assignmentsService.getGrade(assignment.id).subscribe((gradeResponse) => {
                assignment.mark = gradeResponse.mark;
              });
            });
          });
        },
        error: (error) => {
          console.error('Ошибка при создании задания', error);
        }
      });
    }
  }
}