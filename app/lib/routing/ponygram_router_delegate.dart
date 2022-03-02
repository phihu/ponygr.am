import 'package:flutter/material.dart';

import 'ponygram_route_path.dart';
import '../pages/home_page.dart';
import '../pages/search-page.dart';
import '../pages/messages-pages.dart';
import '../pages/post-page.dart';
import '../pages/account_page.dart';

class PonygramRouterDelegate extends RouterDelegate<PonygramRoutePath>
    with ChangeNotifier, PopNavigatorRouterDelegateMixin<PonygramRoutePath> {
  final GlobalKey<NavigatorState> navigatorKey;

  String _page = 'home';
  bool show404 = false;


  PonygramRouterDelegate() : navigatorKey = GlobalKey<NavigatorState>();

  PonygramRoutePath get currentConfiguration {
    if (show404) {
      return PonygramRoutePath.unknown();
    }
//    : PonygramRoutePath.details(books.indexOf(_selectedBook));
    if(_page == 'account') {
      return PonygramRoutePath.account();
    }
    return PonygramRoutePath.home();
  }

  @override
  Widget build(BuildContext context) {
    return Navigator(
      key: navigatorKey,
      pages: [
/*        MaterialPage(
          key: ValueKey('BooksListPage'),
          child: BooksListScreen(
            books: books,
            onTapped: _handleBookTapped,
          ),
        ),
//        if (show404 || _page == 'home')
 */
        if (_page == 'search')
          MaterialPage(key: ValueKey('Search'), child: SearchPage(navigate: _navigate))
        else if (_page == 'messages')
          MaterialPage(key: ValueKey('Messages'), child: MessagesPage(navigate: _navigate))
        else if (_page == 'post')
          MaterialPage(key: ValueKey('Post'), child: PostPage(navigate: _navigate))
        else if (_page == 'account')
          MaterialPage(key: ValueKey('Account'), child: AccountPage(navigate: _navigate))
        else
          MaterialPage(key: ValueKey('Home'), child: HomePage(navigate: _navigate))
      ],
      onPopPage: (route, result) {
        if (!route.didPop(result)) {
          return false;
        }
        // Update the list of pages by setting _selectedBook to null
        show404 = false;
        _page = '';
        notifyListeners();
        return true;
      },
    );
  }

  @override
  Future<void> setNewRoutePath(PonygramRoutePath configuration) async {
    if (configuration.isUnknown) {
      show404 = true;
      return;
    }

/*    if (configuration.isDetailsPage) {
      if (configuration.id < 0 || configuration.id > books.length - 1) {
        show404 = true;
        return;
      }
      _selectedBook = books[configuration.id];
    } else {
      _selectedBook = null;
    }
*/

    show404 = false;
  }

  void _navigate(String page) {
    _page = page;
    notifyListeners();
  }

}
