import {Component, Input} from 'angular2/core'
import {Router} from 'angular2/router';
import {LoginComponent} from '../login/login.component'

@Component({
    selector: 'my-topbar',
    templateUrl: 'app/shared/topbar/topbar.component.html',
    directives: [
        LoginComponent
    ]
})
export class TopbarComponent{
    @Input()
    title: string;

    constructor(
        private _router: Router ) {

    }

    gotoIndex() {
        let link = ['Dashboard'];
        this._router.navigate(link);
        return false;
    }
}
