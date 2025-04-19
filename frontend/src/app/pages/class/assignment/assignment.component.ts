import { Component, OnInit } from '@angular/core';
import { ActivatedRoute } from '@angular/router';
import { HttpClient } from '@angular/common/http';

@Component({
  selector: 'app-assignment',
  templateUrl: './assignment.component.html',
  styleUrls: ['./assignment.component.css'],
})
export class AssignmentComponent implements OnInit {
  assignment: any = null;
  isEditing = false;
  editedStatement: string = '';
  parsedMarkdown: string = '';

  constructor(private route: ActivatedRoute, private http: HttpClient) {}

  ngOnInit(): void {
    const assignmentId = this.route.snapshot.paramMap.get('id');
    this.http
      .get(`http://localhost:8080/classes/assignment/${assignmentId}`, { withCredentials: true })
      .subscribe((data) => {
        this.assignment = data;
        this.editedStatement = this.assignment.statement || '';
        this.parsedMarkdown = (window as any).marked?.parse(this.assignment.statement || '') || '';
      });
  }
  

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
}
