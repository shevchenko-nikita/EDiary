import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { HttpClient } from '@angular/common/http';
import { RoleService } from 'src/app/services/role.service';

@Component({
  selector: 'app-assignment',
  templateUrl: './assignment.component.html',
  styleUrls: ['./assignment.component.css'],
})
export class AssignmentComponent implements OnInit {
  assignment: any = null;
  isEditing = false;
  isEditingTitle = false;
  isEditingDeadline = false;
  editedStatement: string = '';
  editedTitle: string = '';
  editedDeadline: string = '';
  parsedMarkdown: string = '';
  isTeacher = false;

  constructor(
    private route: ActivatedRoute, 
    private router: Router,
    private http: HttpClient,
    private roleService: RoleService
  ) {}

  ngOnInit(): void {
    const assignmentId = this.route.snapshot.paramMap.get('id');
    this.http
      .get(`http://localhost:8080/classes/assignment/${assignmentId}`, { withCredentials: true })
      .subscribe((data) => {
        this.assignment = data;
        this.editedStatement = this.assignment.statement || '';
        this.editedTitle = this.assignment.name || '';
        this.parsedMarkdown = (window as any).marked?.parse(this.assignment.statement || '') || '';
        
        // Format deadline for datetime-local input if exists
        if (this.assignment.deadline) {
          const deadlineDate = new Date(this.assignment.deadline);
          this.editedDeadline = this.formatDateForInput(deadlineDate);
        }
        
        console.log(this.assignment);

        this.roleService.isTeacher(this.assignment.class_id).subscribe(res => {
          this.isTeacher = res.isTeacher;
        });
      });
  }
  
  // Helper method to format date for datetime-local input
  formatDateForInput(date: Date): string {
    const year = date.getFullYear();
    const month = (date.getMonth() + 1).toString().padStart(2, '0');
    const day = date.getDate().toString().padStart(2, '0');
    const hours = date.getHours().toString().padStart(2, '0');
    const minutes = date.getMinutes().toString().padStart(2, '0');
    
    return `${year}-${month}-${day}T${hours}:${minutes}`;
  }

  // Edit task description
  editDescription() {
    this.isEditing = true;
  }

  saveDescription() {
    const assignmentId = this.route.snapshot.paramMap.get('id');
    const payload = {
      ...this.assignment,
      statement: this.editedStatement,
    };

    this.http
      .put(`http://localhost:8080/classes/update-assignment`, payload, { withCredentials: true })
      .subscribe(() => {
        this.assignment.statement = this.editedStatement;
        this.updateMarkdown();
        this.isEditing = false;
      });
  }

  updateMarkdown() {
    this.parsedMarkdown = (window as any).marked?.parse(this.editedStatement) || '';
  }

  // Edit task title
  editTitle() {
    this.isEditingTitle = true;
    this.editedTitle = this.assignment.name;
  }

  saveTitle() {
    if (!this.editedTitle.trim()) {
      return; // Don't save empty title
    }
    
    const payload = {
      ...this.assignment,
      name: this.editedTitle,
    };

    this.http
      .put(`http://localhost:8080/classes/update-assignment`, payload, { withCredentials: true })
      .subscribe(() => {
        this.assignment.name = this.editedTitle;
        this.isEditingTitle = false;
      });
  }

  // Edit deadline
  editDeadline() {
    this.isEditingDeadline = true;
    
    if (this.assignment.deadline) {
      const deadlineDate = new Date(this.assignment.deadline);
      this.editedDeadline = this.formatDateForInput(deadlineDate);
    } else {
      // Set default to current time + 1 week if no deadline exists
      const defaultDate = new Date();
      defaultDate.setDate(defaultDate.getDate() + 7);
      this.editedDeadline = this.formatDateForInput(defaultDate);
    }
  }

  saveDeadline() {
    // Explicitly type the deadline variable as string | null
    let deadline: string | null = null;
    
    if (this.editedDeadline) {
      // Format the date as yyyy-mm-dd hh:mm:ss
      const date = new Date(this.editedDeadline);
      
      const year = date.getFullYear();
      const month = (date.getMonth() + 1).toString().padStart(2, '0');
      const day = date.getDate().toString().padStart(2, '0');
      const hours = date.getHours().toString().padStart(2, '0');
      const minutes = date.getMinutes().toString().padStart(2, '0');
      
      // Set seconds to 00 as requested
      deadline = `${year}-${month}-${day} ${hours}:${minutes}:00`;
    }
    
    const payload = {
      ...this.assignment,
      deadline: deadline,
    };

    this.http
      .put(`http://localhost:8080/classes/update-assignment`, payload, { withCredentials: true })
      .subscribe(() => {
        this.assignment.deadline = deadline;
        this.isEditingDeadline = false;
      });
  }

  removeDeadline() {
    const payload = {
      ...this.assignment,
      deadline: null,
    };

    this.http
      .put(`http://localhost:8080/classes/update-assignment`, payload, { withCredentials: true })
      .subscribe(() => {
        this.assignment.deadline = null;
        this.isEditingDeadline = false;
      });
  }
  
  deleteAssignment() {
    if (confirm('Ви впевнені, що хочете видалити це завдання? Ця дія є незворотною.')) {
      const assignmentId = this.assignment.id;
      
      this.http
        .delete(`http://localhost:8080/classes/delete-assignment/${assignmentId}`, { withCredentials: true })
        .subscribe(() => {
          // После успешного удаления вернуться на страницу класса
          this.router.navigate(['/class', this.assignment.class_id]);
        });
    }
  }
}