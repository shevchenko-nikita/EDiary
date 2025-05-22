import { Component, OnInit } from '@angular/core';
import { trigger, transition, style, animate } from '@angular/animations';
import { HttpClient } from '@angular/common/http';
import { ChartConfiguration, ChartData, ChartType } from 'chart.js';

interface StudentInfo {
  overall_average: number;
  student_classes: number;
  teaching_classes: number;
}

interface ClassCard {
  name?: string;
  grade?: number;
}

interface GradeDistribution {
  labels: string[];
  data: number[];
}

interface Statistic {
  student_info: StudentInfo;
  subjects: ClassCard[];
  grade_distribution: GradeDistribution;
}

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
  statistics: Statistic | null = null;
  loading = false;
  error: string | null = null;

  private apiUrl = 'http://localhost:8080/statistic';

  private subjectColors = ['#4070f4', '#ff6384', '#36a2eb', '#ffcd56', '#4bc0c0', '#9966ff', '#ff9f40'];
  private gradeColors = ['#36a2eb', '#4bc0c0', '#ffcd56', '#ff6384', '#9966ff', '#ff9f40'];

  // Chart.js configuration
  public doughnutChartType = 'doughnut' as const;
  public doughnutChartData: ChartData<'doughnut'> = {
    labels: [],
    datasets: [{
      data: [],
      backgroundColor: [],
      borderColor: '#ffffff',
      borderWidth: 3,
      hoverBackgroundColor: [],
      hoverBorderWidth: 4
    }]
  };

  public doughnutChartOptions: ChartConfiguration<'doughnut'>['options'] = {
    responsive: true,
    maintainAspectRatio: false,
    plugins: {
      legend: {
        display: false // Мы создадим свою легенду
      },
      tooltip: {
        enabled: true,
        backgroundColor: 'rgba(0, 0, 0, 0.8)',
        titleColor: '#ffffff',
        bodyColor: '#ffffff',
        borderColor: '#ffffff',
        borderWidth: 1,
        cornerRadius: 8,
        displayColors: true,
        callbacks: {
          label: (context) => {
            const label = context.label || '';
            const value = context.parsed;
            const total = context.dataset.data.reduce((a: number, b: number) => a + b, 0);
            const percentage = ((value / total) * 100).toFixed(1);
            return `${label}: ${value} (${percentage}%)`;
          }
        }
      }
    },
    cutout: '60%', // Размер внутреннего отверстия
    animation: {
      animateRotate: true,
      animateScale: true,
      duration: 1000,
      easing: 'easeOutQuart'
    },
    elements: {
      arc: {
        borderJoinStyle: 'round'
      }
    }
  };

  mockNotifications = [
    { text: 'Новий тест з математики доступний', time: '2 години тому' },
    { text: 'Ви отримали оцінку 5 з програмування', time: '1 день тому' },
    { text: 'Викладач додав новий матеріал', time: '2 дні тому' }
  ];

  constructor(private http: HttpClient) { }

  ngOnInit(): void {
    this.loadStatistics();
  }

  loadStatistics(): void {
    this.loading = true;
    this.error = null;

    this.http.get<Statistic>(this.apiUrl, { withCredentials: true }).subscribe({
      next: (data: Statistic) => {
        this.statistics = data;
        this.updateChartData();
        console.log(data);
        this.loading = false;
        console.log('Статистика загружена:', data);
      },
      error: (error) => {
        this.loading = false;
        this.error = 'Не вдалося завантажити дані. Перевірте з\'єднання з сервером.';
        console.error('Ошибка при загрузке статистики:', error);
      }
    });
  }

  private updateChartData(): void {
    if (this.statistics?.grade_distribution) {
      const { labels, data } = this.statistics.grade_distribution;
      
      this.doughnutChartData = {
        labels: labels,
        datasets: [{
          data: data,
          backgroundColor: labels.map((_, index) => this.getGradeColor(index)),
          borderColor: '#ffffff',
          borderWidth: 3,
          hoverBackgroundColor: labels.map((_, index) => this.getGradeColorHover(index)),
          hoverBorderWidth: 4
        }]
      };
    }
  }

  getPercentage(grade: number): number {
    return Math.min((grade / 100) * 100, 100);
  }

  getSubjectColor(subjectName: string | undefined): string {
    if (!subjectName) return this.subjectColors[0];
    
    const hash = subjectName.split('').reduce((a, b) => {
      a = ((a << 5) - a) + b.charCodeAt(0);
      return a & a;
    }, 0);
    
    const index = Math.abs(hash) % this.subjectColors.length;
    return this.subjectColors[index];
  }

  getGradeColor(index: number): string {
    return this.gradeColors[index % this.gradeColors.length];
  }

  private getGradeColorHover(index: number): string {
    const baseColor = this.getGradeColor(index);
    // Делаем цвет немного светлее при наведении
    return this.lightenColor(baseColor, 20);
  }

  private lightenColor(color: string, percent: number): string {
    const num = parseInt(color.replace("#", ""), 16);
    const amt = Math.round(2.55 * percent);
    const R = (num >> 16) + amt;
    const G = (num >> 8 & 0x00FF) + amt;
    const B = (num & 0x0000FF) + amt;
    return "#" + (0x1000000 + (R < 255 ? R < 1 ? 0 : R : 255) * 0x10000 +
      (G < 255 ? G < 1 ? 0 : G : 255) * 0x100 +
      (B < 255 ? B < 1 ? 0 : B : 255)).toString(16).slice(1);
  }

  // Методы для получения данных легенды
  getLegendData(): Array<{label: string, value: number, color: string, percentage: string}> {
    if (!this.statistics?.grade_distribution) return [];
    
    const { labels, data } = this.statistics.grade_distribution;
    const total = data.reduce((sum, value) => sum + value, 0);
    
    return labels.map((label, index) => ({
      label,
      value: data[index],
      color: this.getGradeColor(index),
      percentage: total > 0 ? ((data[index] / total) * 100).toFixed(1) : '0'
    }));
  }

  get overallAverage(): number {
    return this.statistics?.student_info?.overall_average || 0;
  }

  get studentClasses(): number {
    return this.statistics?.student_info?.student_classes || 0;
  }

  get teachingClasses(): number {
    return this.statistics?.student_info?.teaching_classes || 0;
  }
}