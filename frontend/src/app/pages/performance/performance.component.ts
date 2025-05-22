import { Component, OnInit } from '@angular/core';
import { trigger, transition, style, animate } from '@angular/animations';

@Component({
  selector: 'app-performance',
  templateUrl: './performance.component.html',
  styleUrls: ['./performance.component.css'],
  animations: [
    trigger('fadeIn', [
      transition(':enter', [
        style({ opacity: 0, transform: 'translateY(20px)' }),
        animate('0.5s ease-out', style({ opacity: 1, transform: 'translateY(0)' }))
      ])
    ]),
    trigger('scaleIn', [
      transition(':enter', [
        style({ opacity: 0, transform: 'scale(0.8)' }),
        animate('0.5s 0.2s ease-out', style({ opacity: 1, transform: 'scale(1)' }))
      ])
    ])
  ]
})
export class PerformanceComponent implements OnInit {
  // Данные для графиков и статистики
  studentClasses = 5;
  teachingClasses = 2;
  avgGrades: { subject: string; grade: number; color: string }[] = [
    { subject: 'Математика', grade: 4.5, color: '#4070f4' },
    { subject: 'Физика', grade: 4.2, color: '#ff6384' },
    { subject: 'Программирование', grade: 4.8, color: '#36a2eb' },
    { subject: 'Литература', grade: 3.9, color: '#ffcd56' },
    { subject: 'История', grade: 4.1, color: '#4bc0c0' }
  ];
  
  // Данные для круговой диаграммы распределения оценок
  gradeDistribution = {
    labels: ['Отлично (5)', 'Хорошо (4)', 'Удовлетворительно (3)', 'Неудовлетворительно (2)'],
    data: [65, 20, 10, 5],
    colors: ['#36a2eb', '#4bc0c0', '#ffcd56', '#ff6384']
  };
  
  // Данные для линейного графика прогресса
  progressData = {
    labels: ['Янв', 'Фев', 'Март', 'Апр', 'Май'],
    data: [3.8, 4.0, 4.2, 4.5, 4.7]
  };
  
  // Последние уведомления
  recentNotifications = [
    { text: 'Новый тест по математике доступен', time: '2 часа назад' },
    { text: 'Вы получили оценку 5 по программированию', time: '1 день назад' },
    { text: 'Преподаватель добавил новый материал', time: '2 дня назад' }
  ];
  
  constructor() { }

  ngOnInit(): void {
    this.setupCharts();
  }

  setupCharts(): void {
    // Здесь можно инициализировать графики, если используются сторонние библиотеки
    // Например, Chart.js или похожие
    console.log('Charts initialized');
  }
  
  // Расчет общего среднего балла
  get overallAverage(): number {
    const sum = this.avgGrades.reduce((total, subject) => total + subject.grade, 0);
    return Math.round((sum / this.avgGrades.length) * 100) / 100;
  }
  
  // Расчет процента выполнения для прогресс-баров
  getPercentage(grade: number): number {
    return (grade / 5) * 100;
  }
}