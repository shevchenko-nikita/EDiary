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
		  	
			const baseURL = 'http://localhost:8080/';
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
  }
  
