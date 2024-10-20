import { Component } from '@angular/core';
import { IconDirective } from '@coreui/icons-angular';
import {
  ContainerComponent,
  RowComponent,
  ColComponent,
  TextColorDirective,
  CardComponent,
  CardBodyComponent,
  FormDirective,
  InputGroupComponent,
  InputGroupTextDirective,
  FormControlDirective,
  ButtonDirective,
} from '@coreui/angular';
import { UserService } from '../../../services/user.service';
import { UserCreate } from '../../../types';
import { FormsModule } from '@angular/forms';
import { ToastrService } from 'ngx-toastr';

@Component({
  selector: 'app-create',
  standalone: true,
  imports: [
    FormsModule,
    ContainerComponent,
    RowComponent,
    ColComponent,
    TextColorDirective,
    CardComponent,
    CardBodyComponent,
    FormDirective,
    InputGroupComponent,
    InputGroupTextDirective,
    IconDirective,
    FormControlDirective,
    ButtonDirective,
  ],
  templateUrl: './create.component.html',
})
export class CreateComponent {
  constructor(private userService: UserService, private toast: ToastrService) {}

  user: UserCreate = {
    username: '',
    email: '',
    firstName: '',
    lastName: '',
  };

  onSubmitButton(event: MouseEvent) {
    this.userService.createStaff(this.user).subscribe(
      () => {
        this.toast.success('Create staff successfully');
        this.user = {
          username: '',
          email: '',
          firstName: '',
          lastName: '',
        };
      },
      (err) => {
        console.log(err);
        this.toast.error('Create staff failed', err.error.error);
      }
    );

    event.stopPropagation();
  }
}
