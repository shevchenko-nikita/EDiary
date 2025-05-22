import { Component, OnInit } from '@angular/core';
import { trigger, transition, style, animate } from '@angular/animations';
import { HttpClient } from '@angular/common/http';

// Интерфейсы для типизации данных
interface StudentInfo {
  overall_average: number;
  student_classes: number;
  teaching_classes: number;
}

interface ClassCard {
  name?: string;
  grade?: number;
  // Добавьте другие поля, если они есть в вашей модели ClassCard
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
  // Основные данные
  statistics: Statistic | null = null;
  loading = false;
  error: string | null = null;

  // API URL
  private apiUrl = 'http://localhost:8080/statistic';

  // Цвета для предметов и графиков
  private subjectColors = ['#4070f4', '#ff6384', '#36a2eb', '#ffcd56', '#4bc0c0', '#9966ff', '#ff9f40'];
  private gradeColors = ['#36a2eb', '#4bc0c0', '#ffcd56', '#ff6384'];

  // Временные данные для уведомлений (пока не подключены к API)
  mockNotifications = [
    { text: 'Новий тест з математики доступний', time: '2 години тому' },
    { text: 'Ви отримали оцінку 5 з програмування', time: '1 день тому' },
    { text: 'Викладач додав новий матеріал', time: '2 дні тому' }
  ];

  constructor(private http: HttpClient) { }

  ngOnInit(): void {
    this.loadStatistics();
  }

  /**
   * Загрузка статистики с backend
   */
  loadStatistics(): void {
    this.loading = true;
    this.error = null;

    this.http.get<Statistic>(this.apiUrl, { withCredentials: true }).subscribe({
      next: (data: Statistic) => {
        this.statistics = data;
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

  /**
   * Расчет процента для прогресс-баров предметов
   */
  getPercentage(grade: number): number {
    return Math.min((grade / 100) * 100, 100);
  }

  /**
   * Получение цвета для предмета по его названию
   */
  getSubjectColor(subjectName: string | undefined): string {
    if (!subjectName) return this.subjectColors[0];
    
    // Простой алгоритм для присвоения цвета на основе названия
    const hash = subjectName.split('').reduce((a, b) => {
      a = ((a << 5) - a) + b.charCodeAt(0);
      return a & a;
    }, 0);
    
    const index = Math.abs(hash) % this.subjectColors.length;
    return this.subjectColors[index];
  }

  /**
   * Получение цвета для элемента распределения оценок
   */
  getGradeColor(index: number): string {
    return this.gradeColors[index % this.gradeColors.length];
  }

  /**
   * Расчет трансформации для сегментов круговой диаграммы
   * Упрощенная версия - для более точной диаграммы рекомендуется использовать Chart.js
   */
  getSegmentTransform(index: number, data: number[]): string {
    const total = data.reduce((sum, value) => sum + value, 0);
    if (total === 0) return 'rotate(0deg) skew(0deg)';

    let currentAngle = 0;
    for (let i = 0; i < index; i++) {
      currentAngle += (data[i] / total) * 360;
    }

    const segmentAngle = (data[index] / total) * 360;
    const skewAngle = 90 - segmentAngle;

    return `rotate(${currentAngle}deg) skew(${Math.max(skewAngle, 0)}deg)`;
  }

  /**
   * Получение общего среднего балла (для совместимости, если нужно использовать в других местах)
   */
  get overallAverage(): number {
    return this.statistics?.student_info?.overall_average || 0;
  }

  /**
   * Получение количества учебных классов
   */
  get studentClasses(): number {
    return this.statistics?.student_info?.student_classes || 0;
  }

  /**
   * Получение количества преподаваемых классов
   */
  get teachingClasses(): number {
    return this.statistics?.student_info?.teaching_classes || 0;
  }
}