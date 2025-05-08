import { Component, OnInit } from '@angular/core';
import { FormBuilder, FormGroup } from '@angular/forms';
import { HttpClient } from '@angular/common/http';

@Component({
	selector: 'app-profile',
	templateUrl: './profile.component.html',
	styleUrls: ['./profile.component.css']
})

export class ProfileComponent {
	profileForm: FormGroup;
	profileImage: string = '';
  
	constructor(private fb: FormBuilder, private http: HttpClient) {
	  this.profileForm = this.fb.group({
		first_name: [''],
		middle_name: [''],
		second_name: [''],
		username: [''],
		email: ['']
	  });
	}

	ngOnInit() {
		this.loadProfile();
	}

	loadProfile() {
		this.http.get<any>('http://localhost:8080/user/profile', { withCredentials: true }).subscribe(data => {
			this.profileForm.patchValue({
				first_name: data.first_name,
				middle_name: data.middle_name,
				second_name: data.second_name,
				username: data.username,
				email: data.email
		  	});
		  	
			const baseURL = 'http://localhost:8080/images/';
			this.profileImage = baseURL + data.profile_img_path;
			console.log(this.profileImage);
		});
	}
  
	onSubmit() {
		this.http.put('http://localhost:8080/user/update-profile', this.profileForm.value, { withCredentials: true }).subscribe(response => {
			console.log('Профіль оновлено', response);
		});
	}
  
	changeImage() {
	  alert('Функція зміни зображення ще не реалізована');
	}

	onFileSelected(event: Event) {
		const input = event.target as HTMLInputElement;
	
		if (input.files && input.files[0]) {
			const file = input.files[0];
			const formData = new FormData();
			console.log(file);
			formData.append('profile_image', file); // adjust key if backend expects a different name
	
			this.http.put('http://localhost:8080/user/update-profile-image', formData, {
				withCredentials: true
			}).subscribe({
				next: (response: any) => {
					console.log('Зображення оновлено', response);
					this.loadProfile(); // reload profile to update image preview
				},
				error: err => {
					console.error('Помилка під час оновлення зображення', err);
				}
			});
		}
	}
	
  }
  
