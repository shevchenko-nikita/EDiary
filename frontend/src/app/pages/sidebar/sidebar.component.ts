import { Component, ElementRef, ViewChild, AfterViewInit, Renderer2 } from '@angular/core';

@Component({
	selector: 'app-sidebar',
	templateUrl: './sidebar.component.html',
	styleUrls: ['./sidebar.component.css']
})
export class SidebarComponent implements AfterViewInit {
	@ViewChild('sidebar') sidebar!: ElementRef;

	constructor(private renderer: Renderer2) {}

	ngAfterViewInit(): void {
		this.renderer.addClass(this.sidebar.nativeElement, 'close'); // Закрываем по умолчанию
	}

	hideSidebar(): void {
		this.renderer.addClass(this.sidebar.nativeElement, 'close');
	}

	showSidebar(): void {
		this.renderer.removeClass(this.sidebar.nativeElement, 'close');
	}
}
