import { Component, OnInit } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { ActivatedRoute } from '@angular/router';
import { RoleService } from 'src/app/services/role.service';

@Component({
  selector: 'app-users',
  templateUrl: './users.component.html',
  styleUrls: ['./users.component.css']
})
export class UsersComponent implements OnInit {
  students: any[] = [];
  teacher: any = {};
  classID: number = 1;
  class_code: string = '';
  showDeleteModal = false;
  showClassCodeModal = false;
  studentToDelete: string | null = null;
  isTeacher = false;

  constructor(
    private http: HttpClient, 
    private route: ActivatedRoute,
    private roleService: RoleService) { }

  ngOnInit(): void {
    this.route.paramMap.subscribe(params => {
      this.classID = Number(params.get('id'));
      if (this.classID) {
        this.loadData();
      }
      this.roleService.isTeacher(this.classID).subscribe(res => {
        this.isTeacher = res.isTeacher;
      });
    });
  }

  loadData(): void {
    this.http.get<any[]>(`http://localhost:8080/classes/student-list/${this.classID}`, { withCredentials: true })
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

    this.http.get<any>(`http://localhost:8080/classes/teacher/${this.classID}`, { withCredentials: true })
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

      this.http.get<any>(`http://localhost:8080/classes/get-info/${this.classID}`, { withCredentials: true })
      .subscribe(
        (response) => {
          this.class_code = response.class_code
        },
        (error) => {
          console.error('Error fetching teacher data:', error);
        }
      );
  }

  openDeleteModal(studentID: string): void {
    this.studentToDelete = studentID;
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
