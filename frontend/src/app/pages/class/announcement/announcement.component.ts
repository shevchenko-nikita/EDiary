import { Component, OnInit } from '@angular/core';
import { ActivatedRoute } from '@angular/router';
import { HttpClient } from '@angular/common/http';

interface Announcement {
    id: number;
    user_name: string;
    time_posted: string;
    user_profile_img: string;
    text: string;
    editing?: boolean;
    showMenu?: boolean;
}

interface NewAnnouncement {
  class_id: number;
  text: string;
}

@Component({
    selector: 'app-announcement',
    templateUrl: './announcement.component.html',
    styleUrls: ['./announcement.component.css']
})
export class AnnouncementComponent implements OnInit {
    classId!: number;
    subjectName: string = '';
    newMessage: string = '';
    messages: Announcement[] = [];

    constructor(private route: ActivatedRoute, private http: HttpClient) {}

    ngOnInit() {
        this.route.paramMap.subscribe(params => {
            this.classId = Number(params.get('id'));
            this.loadMessages();
            this.loadClassInfo();
        });
    }

    loadClassInfo() {
      this.http.get<{ name: string }>(`http://localhost:8080/classes/get-info/${this.classId}`, { withCredentials: true })
          .subscribe(data => {
              this.subjectName = data.name;
              console.log(data);
          }, error => {
              console.error('Помилка при завантаженні інформації про клас:', error);
          });
    }

    loadMessages() {
        this.http.get<Announcement[]>(`http://localhost:8080/classes/all-messages/${this.classId}`, { withCredentials: true })
            .subscribe(data => {
                this.messages = data.map(cls => ({
                  ...cls,
                  user_profile_img: 'http://localhost:8080/' + cls.user_profile_img
                })).reverse();
                // console.log(this.messages);
            }, error => {
                console.error('Помилка при завантаженні повідомлень:', error);
            });
    }

    sendMessage() {
        if (!this.newMessage.trim()) return;

        const newAnnouncement: NewAnnouncement = {
            class_id: this.classId,
            text: this.newMessage,
        };

        this.http.post('http://localhost:8080/classes/create-message', {
            class_id: this.classId,
            text: this.newMessage
        }, { withCredentials: true }).subscribe(() => {
            // this.messages.unshift(newAnnouncement);
            this.newMessage = '';
        }, error => {
            console.error('Помилка при відправленні повідомлення:', error);
        });
    }

    toggleMenu(message: Announcement) {
        message.showMenu = !message.showMenu;
    }

    editMessage(message: Announcement) {
        message.editing = true;
    }

    saveEdit(message: Announcement) {
        message.editing = false;
        // Логика для сохранения изменений
    }

    deleteMessage(message: Announcement) {
        this.messages = this.messages.filter(m => m.id !== message.id);
        // Добавить логику удаления на сервере
    }

    onMessageInput() {
        // Дополнительная логика, если нужно
    }
}
