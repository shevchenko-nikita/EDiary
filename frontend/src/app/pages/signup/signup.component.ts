import {
	Component,
	AfterViewInit,
	ElementRef,
	QueryList,
	ViewChildren,
} from '@angular/core';

import { HttpClient } from '@angular/common/http';
import { Router } from '@angular/router';

@Component({
	selector: 'app-signup',
	templateUrl: './signup.component.html',
	styleUrls: ['./signup.component.css'],
})
export class SignupComponent implements AfterViewInit {
	currentStep = 0;

	@ViewChildren('page') pages!: QueryList<ElementRef>;

	firstName = '';
	middleName = '';
	secondName = '';
	email = '';
	username = '';
	password = '';
	confirmPassword = '';
	errorMessage = '';

	constructor(private http: HttpClient, private router: Router) {}

	ngAfterViewInit(): void {
		this.showPage(this.currentStep);
	}

	showPage(pageIndex: number) {
		this.pages.forEach((page, index) => {
			(page.nativeElement as HTMLElement).style.display =
				index === pageIndex ? 'block' : 'none';
		});
		this.currentStep = pageIndex;
	}

	nextPage() {
		if (this.currentStep < this.pages.length - 1) {
			this.showPage(this.currentStep + 1);
		}
	}

	prevPage() {
		if (this.currentStep > 0) {
			this.showPage(this.currentStep - 1);
		}
	}

	onSubmit() 
	{
		if (this.password != this.confirmPassword)
    	{
			console.error('Error: ', 'Паролі не співпадають');
			this.errorMessage = 'Паролі не співпадають!';
			return;
		}

		const userData = {
			first_name: this.firstName,
			middle_name: this.middleName,
			second_name: this.secondName,
			email: this.email,
			username: this.username,
			password: this.password
		}

		this.http.post('http://localhost:8080/sign-up', userData).subscribe({
			next: () => {
				this.router.navigate(['/log-in']);
			},
			error: (err) => {
				console.error('Error: ', err);
				this.errorMessage = 'Помилка реєстрації! Спробуйте ще раз.';
			},
		});
	}
}
