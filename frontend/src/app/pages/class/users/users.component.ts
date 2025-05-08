import { Component, OnInit } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { ActivatedRoute } from '@angular/router';

@Component({
  selector: 'app-users',
  templateUrl: './users.component.html',
  styleUrls: ['./users.component.css']
})
export class UsersComponent implements OnInit {
  students: any[] = [];
  teacher: any = {};
  classId: number = 1;
  class_code: string = '';
  showDeleteModal = false;
  showClassCodeModal = false;
  studentToDelete: string | null = null;

  constructor(private http: HttpClient, private route: ActivatedRoute) { }

  ngOnInit(): void {
    this.route.paramMap.subscribe(params => {
      this.classId = Number(params.get('id'));
      if (this.classId) {
        this.loadData();
      }
    });
  }

  loadData(): void {
    this.http.get<any[]>(`http://localhost:8080/classes/student-list/${this.classId}`, { withCredentials: true })
      .subscribe(
        (data) => {
          this.students = data.map(cls => ({
            ...cls,
            profile_img_path: 'http://localhost:8080/images/' + cls.profile_img_path
          }));

          console.log(data);
        },
        (error) => {
          console.error('Error fetching students data:', error);
        }
      );

    this.http.get<any>(`http://localhost:8080/classes/teacher/${this.classId}`, { withCredentials: true })
      .subscribe(
        (response) => {
          this.teacher = response;
          this.teacher.profile_img_path = 'http://localhost:8080/images/' + this.teacher.profile_img_path 
          // this.teacher.username = '@' + this.teacher.username
        },
        (error) => {
          console.error('Error fetching teacher data:', error);
        }
      );

      this.http.get<any>(`http://localhost:8080/classes/get-info/${this.classId}`, { withCredentials: true })
      .subscribe(
        (response) => {
          this.class_code = response.class_code
        },
        (error) => {
          console.error('Error fetching teacher data:', error);
        }
      );
  }

  openDeleteModal(studentId: string): void {
    this.studentToDelete = studentId;
    this.showDeleteModal = true;
  }

  closeDeleteModal(): void {
    this.showDeleteModal = false;
    this.studentToDelete = null;
  }

  confirmDelete(): void {
    if (this.studentToDelete) {
      this.students = this.students.filter(student => student.id !== this.studentToDelete);
      this.closeDeleteModal();
    }
  }

  openClassCodeModal(): void {
    this.showClassCodeModal = true;
  }

  closeClassCodeModal(): void {
    this.showClassCodeModal = false;
  }
}
