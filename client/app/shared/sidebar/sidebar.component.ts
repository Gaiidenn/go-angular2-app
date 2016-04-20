import {Component} from 'angular2/core';
import {Router} from 'angular2/router';

@Component({
    selector: 'my-sidebar',
    templateUrl: 'app/shared/sidebar/sidebar.component.html'
})
export class SidebarComponent {
    constructor(
        private _router: Router ) {

    }

    isActive(name: string) {
        let instruction = this._router.generate([name]);
        return this._router.isRouteActive(instruction);
    }

    goTo(page: string) {
        let link = [page];
        this._router.navigate(link);
        return false;
    }
}
