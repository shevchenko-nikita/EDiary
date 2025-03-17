import { Component, OnInit } from '@angular/core';
import { HttpClient } from '@angular/common/http';

interface Class {
  class_id: number;
  class_code: string;
  name: string;
  teacher_name: string;
  profile_img_path: string;
  grade: number;
}

@Component({
  selector: 'app-education',
  templateUrl: './education.component.html',
  styleUrls: ['./education.component.css']
})
export class EducationComponent implements OnInit {
  classes: Class[] = [];
  isJoinModalOpen = false;
  isExitModalOpen = false;
  classCode: string = '';
  selectedClassId: number | null = null;

  constructor(private http: HttpClient) {}

  ngOnInit() {
    this.loadClasses();
  }

  loadClasses() {
    this.http.get<Class[]>('http://localhost:8080/classes/education-list', { withCredentials: true })
      .subscribe(data => {
        this.classes = data.map(cls => ({
          ...cls,
          profile_img_path: 'http://localhost:8080/' + cls.profile_img_path
        }));
        console.log(this.classes);
      });
  }
  

  openJoinModal() {
    this.isJoinModalOpen = true;
  }

  closeJoinModal() {
    this.isJoinModalOpen = false;
    this.classCode = '';
  }

  joinClass() {
    if (!this.classCode.trim()) return;
    
    this.http.post(`http://localhost:8080/join-class/${this.classCode}`, {})
      .subscribe(() => {
        this.closeJoinModal();
        this.loadClasses();
      });
  }

  leaveClass(classId: number) {
    this.selectedClassId = classId;
    this.isExitModalOpen = true;
  }

  closeExitModal() {
    this.isExitModalOpen = false;
    this.selectedClassId = null;
  }

  confirmLeaveClass() {
    if (this.selectedClassId !== null) {
      this.http.delete(`http://localhost:8080/classes/leave-class/${this.selectedClassId}`, { withCredentials: true })
        .subscribe(() => {
          this.closeExitModal();
          this.loadClasses();
        });
    }
  }
}
