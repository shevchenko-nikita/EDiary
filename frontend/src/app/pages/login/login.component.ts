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

	constructor(private http: HttpClient, private router: Router) {}

	onLogin() {
		this.errorMessage = '';
		console.log('Вхідні дані:', { username: this.username, password: this.password });
		const loginData = { username: this.username, password: this.password };

		this.http.post('http://localhost:8080/sign-in', loginData, { withCredentials: true }).subscribe({
			next: () => {
				console.log('Авторизація успішна!');
				this.router.navigate(['']);
			},
			error: (err) => {
				console.error('Помилка авторизації:', err);
				this.errorMessage = 'Невірний username або пароль!';
			},
		});
	}

}
