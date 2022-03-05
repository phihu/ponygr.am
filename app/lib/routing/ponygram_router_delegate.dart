import 'package:flutter/material.dart';

import 'ponygram_route_path.dart';
import '../pages/home.dart';
import '../pages/search.dart';
import '../pages/messages.dart';
import '../pages/new_post.dart';
import '../pages/account.dart';

class PonygramRouterDelegate extends RouterDelegate<PonygramRoutePath>
    with ChangeNotifier, PopNavigatorRouterDelegateMixin<PonygramRoutePath> {
  @override
  final GlobalKey<NavigatorState> navigatorKey;

  String _page = 'home';
  bool show404 = false;


  PonygramRouterDelegate() : navigatorKey = GlobalKey<NavigatorState>();

  @override
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
          MaterialPage(key: const ValueKey('Search'), child: SearchPage(navigate: navigate))
        else if (_page == 'messages')
          MaterialPage(key: const ValueKey('Messages'), child: MessagesPage(navigate: navigate))
        else if (_page == 'new')
          MaterialPage(key: const ValueKey('New'), child: NewPostPage(navigate: navigate))
        else if (_page == 'account')
          MaterialPage(key: const ValueKey('Account'), child: AccountPage(navigate: navigate))
        else
          MaterialPage(key: const ValueKey('Home'), child: HomePage(navigate: navigate))
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

  Future<bool> navigate(String page) {
    _page = page;
    notifyListeners();
    return Future.value(true);
  }
}
