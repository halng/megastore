import { Component } from '@angular/core';
import { signal } from '@angular/core';
import { MatButtonModule } from '@angular/material/button';
import { MatFormFieldModule } from '@angular/material/form-field';
import { MatIconModule } from '@angular/material/icon';
import { MatInputModule } from '@angular/material/input';
import { FormsModule, ReactiveFormsModule } from '@angular/forms';
import { UserService } from '../../services/user.service';
@Component({
  selector: 'app-login',
  standalone: true,
  imports: [
    ReactiveFormsModule,
    MatButtonModule,
    MatFormFieldModule,
    MatIconModule,
    MatInputModule,
    FormsModule,
  ],
  templateUrl: './login.component.html',
  styleUrl: './login.component.scss',
  providers: [],
})
export class LoginComponent {
  hide = signal(true);
  username = '';
  password = '';
  error = '';

  constructor(private userService: UserService) {}

  onSubmitForm(event: MouseEvent) {
    // event.preventDefault();
    console.log('Username:', this.username, this.password);
    // check validate username and password
    if (this.username.length < 6 && this.password.length < 8) {
      this.error = 'Username or password is incorrect format';
      // TODO: add alert to show error
      this.username = '';
      this.password = '';
    } else {
      // send api to server
      this.userService.login(this.username, this.password).subscribe(
        (res: any) => {
          console.log('res', res);
          // res -> data -> api-token
          const data = res.data;
          if (!data || !data['api-token']) {
            this.error = 'There was an error while login. Please try again';
            return;
          } else {
            this.userService.setApiToken(data['api-token']);
          }
        },
        (err) => {
          this.error = err.error.error;
        }
      );
    }

    event.stopPropagation();
  }

  onClickHidePassword(event: MouseEvent) {
    // event.preventDefault();
    this.hide.set(!this.hide());
    event.stopPropagation();
  }
}
