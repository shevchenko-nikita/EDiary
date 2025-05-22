import { Component } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { Router } from '@angular/router';

@Component({
	selector: 'app-login',
	templateUrl: './login.component.html',
	styleUrls: ['./login.component.css'],
})

export class LoginComponent 
{
	username = '';
	password = '';
	errorMessage = '';
	showPassword = false;

	constructor(private http: HttpClient, private router: Router) {}

	togglePasswordVisibility() {
		this.showPassword = !this.showPassword;
	}

	onLogin() {
		this.errorMessage = '';
		const loginData = { username: this.username, password: this.password };

		this.http.post('http://localhost:8080/sign-in', loginData, { withCredentials: true }).subscribe({
			next: () => {
				this.router.navigate(['']);
			},
			error: (err) => {
				this.errorMessage = 'Невірний username або пароль!';
			},
		});
	}

}
