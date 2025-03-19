import { Component, OnInit } from '@angular/core';
import { ActivatedRoute } from '@angular/router';

@Component({
  selector: 'app-topbar',
  templateUrl: './topbar.component.html',
  styleUrls: ['./topbar.component.css']
})
export class TopbarComponent implements OnInit {
  classId!: number;

  constructor(private route: ActivatedRoute) {}

  ngOnInit() {
      this.route.paramMap.subscribe(params => {
          this.classId = Number(params.get('id'));
      });
  }
}
