import { Component } from '@angular/core';
import { RouterLink, RouterOutlet } from '@angular/router';
import { NgScrollbar } from 'ngx-scrollbar';

import { IconDirective } from '@coreui/icons-angular';
import {
  ContainerComponent,
  ShadowOnScrollDirective,
  SidebarBrandComponent,
  SidebarComponent,
  SidebarFooterComponent,
  SidebarHeaderComponent,
  SidebarNavComponent,
  SidebarToggleDirective,
  SidebarTogglerDirective,
} from '@coreui/angular';

import { navItems } from './nav';
import { CustomFooterComponent } from './footer/footer.component';
import { CustomHeaderComponent } from './header/header.component';

// function isOverflown(element: HTMLElement) {
//   return (
//     element.scrollHeight > element.clientHeight ||
//     element.scrollWidth > element.clientWidth
//   );
// }

@Component({
  selector: 'app-dashboard',
  templateUrl: './layout.component.html',
  styleUrls: ['./layout.component.scss'],
  standalone: true,
  imports: [
    SidebarComponent,
    SidebarHeaderComponent,
    SidebarBrandComponent,
    RouterLink,
    IconDirective,
    NgScrollbar,
    SidebarNavComponent,
    SidebarFooterComponent,
    SidebarToggleDirective,
    SidebarTogglerDirective,
    CustomFooterComponent,
    ShadowOnScrollDirective,
    ContainerComponent,
    RouterOutlet,
    CustomHeaderComponent,
  ],
})
export class LayoutComponent {
  public navItems = navItems;

  onScrollbarUpdate($event: any) {
    if ($event.verticalUsed) {
      console.log('verticalUsed', $event.verticalUsed);
    }
  }
}
