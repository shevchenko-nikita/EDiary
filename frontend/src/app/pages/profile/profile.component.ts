import { Component } from '@angular/core';
import { FormBuilder, FormGroup } from '@angular/forms';

@Component({
	selector: 'app-profile',
	templateUrl: './profile.component.html',
	styleUrls: ['./profile.component.css']
})

export class ProfileComponent {
	profileForm: FormGroup;
	profileImage: string = '';
  
	constructor(private fb: FormBuilder) {
	  this.profileForm = this.fb.group({
		firstName: [''],
		fatherName: [''],
		lastName: [''],
		username: [''],
		email: ['']
	  });
	}
  
	onSubmit() {
	  console.log(this.profileForm.value);
	}
  
	changeImage() {
	  alert('Функція зміни зображення ще не реалізована');
	}
  }
  
