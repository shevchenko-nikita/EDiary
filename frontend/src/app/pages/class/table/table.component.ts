import { Component, OnInit } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { ActivatedRoute } from '@angular/router';
import { RoleService } from 'src/app/services/role.service';

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
	id?: number;
	student_id: number;
	assignment_id: number;
	mark: number;
	class_id?: number;
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
	isTeacher: boolean = false;
	editMode: { studentId: number, assignmentId: number } | null = null;
	currentMark: number = 0;

	constructor(
		private http: HttpClient, 
		private route: ActivatedRoute,
		private roleService: RoleService
	) {}

	ngOnInit() {
		this.route.paramMap.subscribe(params => {
			const id = params.get('id'); 
			console.log(id);
			if (id) {
				this.classId = +id; 
				this.loadTableData();
				this.checkTeacherRole();
			}
		});
	}

	checkTeacherRole() {
		this.roleService.isTeacher(this.classId).subscribe(response => {
			this.isTeacher = response.isTeacher;
			console.log('Is teacher:', this.isTeacher);
		});
	}

	loadTableData() {
		this.http.get<any>(`http://localhost:8080/classes/table/${this.classId}`, { withCredentials: true })
		.subscribe(response => {
			this.assignments = response.assignments;
			this.students = response.students.sort((a: Student, b: Student) => {
				const lastNameCompare = a.second_name.localeCompare(b.second_name);
				if (lastNameCompare !== 0) return lastNameCompare;
				return a.first_name.localeCompare(b.first_name);
			});
			this.marks = response.marks;

			// console.log(this.assignments);
			// console.log(this.students);
			// console.log(this.marks);
			// console.log(this.classId);
		});
	}

	getMark(studentId: number, assignmentId: number): number {
		if(!this.marks) {
			return 0;
		}

		const mark = this.marks.find(m => m.student_id === studentId && m.assignment_id === assignmentId);
		return mark ? mark.mark : 0;
	}

	getMarkObject(studentId: number, assignmentId: number): Mark | undefined {
		if(!this.marks) {
			return undefined;
		}
		return this.marks.find(m => m.student_id === studentId && m.assignment_id === assignmentId);
	}

	getTotalMark(studentId: number): number {
		return this.assignments.reduce((sum, a) => sum + this.getMark(studentId, a.id), 0);
	}

	startEditingMark(studentId: number, assignmentId: number) {
		if (!this.isTeacher) return;
		
		this.editMode = { studentId, assignmentId };
		this.currentMark = this.getMark(studentId, assignmentId);
	}

	cancelEditing() {
		this.editMode = null;
	}

	saveMark() {
		if (!this.editMode) return;

		const { studentId, assignmentId } = this.editMode;
		const markObject = this.getMarkObject(studentId, assignmentId);
		
		const updatedMark: Mark = {
			id: markObject?.id,
			student_id: studentId,
			assignment_id: assignmentId,
			mark: this.currentMark,
			class_id: this.classId
		};

		this.http.put('http://localhost:8080/classes/grade-assignment', updatedMark, { withCredentials: true })
			.subscribe({
				next: (response) => {
					console.log('Mark updated successfully', response);
					
					// Update local data
					if (markObject) {
						markObject.mark = this.currentMark;
					} else {
						this.marks.push({
							student_id: studentId,
							assignment_id: assignmentId,
							mark: this.currentMark
						});
					}
					
					this.editMode = null;
				},
				error: (error) => {
					console.error('Error updating mark', error);
					this.editMode = null;
				}
			});
	}
}