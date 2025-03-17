import { Component, OnInit } from '@angular/core';
import { HttpClient } from '@angular/common/http';

interface Class {
  id: number;
  class_code: string;
  name: string;
  teacher_id: number;
}

@Component({
  selector: 'app-teaching',
  templateUrl: './teaching.component.html',
  styleUrls: ['./teaching.component.css']
})
export class TeachingComponent implements OnInit {
  classes: Class[] = [];
  isModalOpen = false;
  newClassName = '';

  constructor(private http: HttpClient) {}

  ngOnInit(): void {
    this.fetchClasses();
  }

  fetchClasses(): void {
    this.http.get<Class[]>('http://localhost:8080/classes/teaching-list', { withCredentials: true })
      .subscribe(data => {
        this.classes = data;
      }, error => {
        console.error('Помилка при отриманні списку класів', error);
      });
  }

  openModal(): void {
    this.isModalOpen = true;
  }

  closeModal(): void {
    this.isModalOpen = false;
    this.newClassName = '';
  }

  createClass(): void {
    if (!this.newClassName.trim()) return;

    const payload = { class_name: this.newClassName };

    console.log(payload.class_name);

    this.http.post('http://localhost:8080/classes/create-new-class', payload, { withCredentials: true })
      .subscribe(response => {
        this.fetchClasses();  // Обновить список классов
        this.closeModal();
      }, error => {
        console.error('Помилка при створенні класу', error);
      });
  }
}
