import { Component, OnInit } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { ActivatedRoute } from '@angular/router';

interface Assignment {
	id: number;
	name: string;
}

interface Student {
	id: number;
	first_name: string;
	second_name: string;
	middle_name: string;
}

interface Mark {
	student_id: number;
	assignment_id: number;
	mark: number;
}

@Component({
	selector: 'app-table',
	templateUrl: './table.component.html',
	styleUrls: ['./table.component.css']
})
export class TableComponent implements OnInit {
	assignments: Assignment[] = [];
	students: Student[] = [];
	marks: Mark[] = [];
	classId!: number; 

	constructor(private http: HttpClient, private route: ActivatedRoute) {}

	ngOnInit() {

		this.route.paramMap.subscribe(params => {
			const id = params.get('class_id'); 
			if (id) {
				this.classId = +id; 
				this.loadTableData();
			}
		});
	}

	loadTableData() {
		this.http.get<any>(`http://localhost:8080/classes/table/${this.classId}`, { withCredentials: true })
		.subscribe(response => {
			this.assignments = response.assignments;
			this.students = response.students;
			this.marks = response.marks;
			console.log(this.assignments);
			console.log(this.students);
			console.log(this.marks);
			console.log(this.classId);
		});
	}

	getMark(studentId: number, assignmentId: number): number {
		const mark = this.marks.find(m => m.student_id === studentId && m.assignment_id === assignmentId);
		return mark ? mark.mark : 0;
	}

	getTotalMark(studentId: number): number {
		return this.assignments.reduce((sum, a) => sum + this.getMark(studentId, a.id), 0);
	}
}
