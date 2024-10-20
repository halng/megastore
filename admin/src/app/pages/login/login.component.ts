import { Component } from '@angular/core';
import { UserService } from '../../services/user.service';
import { Router } from '@angular/router';
import { NgStyle } from '@angular/common';
import { IconDirective } from '@coreui/icons-angular';
import {
  ContainerComponent,
  RowComponent,
  ColComponent,
  CardGroupComponent,
  TextColorDirective,
  CardComponent,
  CardBodyComponent,
  FormDirective,
  InputGroupComponent,
  InputGroupTextDirective,
  FormControlDirective,
  ButtonDirective,
} from '@coreui/angular';
import { FormsModule } from '@angular/forms';
import { ToastrService } from 'ngx-toastr';

@Component({
  selector: 'app-login',
  standalone: true,
  imports: [
    FormsModule,
    ContainerComponent,
    RowComponent,
    ColComponent,
    CardGroupComponent,
    TextColorDirective,
    CardComponent,
    CardBodyComponent,
    FormDirective,
    InputGroupComponent,
    InputGroupTextDirective,
    IconDirective,
    FormControlDirective,
    ButtonDirective,
    NgStyle,
  ],
  templateUrl: './login.component.html',
  providers: [],
})
export class LoginComponent {
  username = '';
  password = '';

  constructor(
    private userService: UserService,
    private router: Router,
    private toast: ToastrService
  ) {}

  onSubmitForm(event: MouseEvent) {
    if (this.username.length < 6 && this.password.length < 8) {
      this.toast.warning('Username or password is incorrect format');
      this.username = '';
      this.password = '';
    } else {
      // send api to server
      this.userService.login(this.username, this.password).subscribe(
        (res: any) => {
          // res -> data -> api-token
          const data = res.data;
          if (!data) {
            this.toast.error(
              'There was an error while login. Please try again'
            );
            return;
          } else {
            this.toast.success('Login successfully. Redirecting...');
            this.userService.setApiToken(data);
            // redirect to home page
            this.router.navigate(['/home']);
          }
        },
        (err) => {
          console.log(err)
          const error = err.error.error;
          this.toast.error(error);
        }
      );
    }

    event.stopPropagation();
  }
}
