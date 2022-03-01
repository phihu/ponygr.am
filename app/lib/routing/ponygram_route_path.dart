class PonygramRoutePath {
//  final int id;
  final String page;
  final bool isUnknown;

  PonygramRoutePath.home()
      : page = 'home',
        isUnknown = false;

  PonygramRoutePath.account()
      : page = 'account',
        isUnknown = false;

//  PonygramRoutePath.details(this.id) : isUnknown = false;

  PonygramRoutePath.unknown()
      : page = '404',
        isUnknown = true;

  bool get isHomePage => page == 'home';
  bool get isAccountPage => page == 'account';

//  bool get isDetailsPage => id != null;
}