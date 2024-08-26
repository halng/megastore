import { Component } from '@angular/core';
import { RouterOutlet } from '@angular/router';
import { Title } from '@angular/platform-browser';

@Component({
  selector: 'app-root',
  standalone: true,
  imports: [RouterOutlet],
  templateUrl: './app.component.html',
  styleUrl: './app.component.scss'
})
export class AppComponent {
  tabTitle = 'Mega Store';
  title = 'Welcome to Mega Store';
  constructor(private titleService: Title) {
    this.setTitle(this.tabTitle);
  }

  public setTitle(newTitle: string) {
    this.titleService.setTitle(newTitle);
  }
}
