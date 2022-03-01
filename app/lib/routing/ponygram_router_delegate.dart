import 'package:flutter/material.dart';

import 'ponygram_route_path.dart';
import '../pages/home_page.dart';
import '../pages/account_page.dart';

class PonygramRouterDelegate extends RouterDelegate<PonygramRoutePath>
    with ChangeNotifier, PopNavigatorRouterDelegateMixin<PonygramRoutePath> {
  final GlobalKey<NavigatorState> navigatorKey;

  String page = 'home';
  bool show404 = false;


  PonygramRouterDelegate() : navigatorKey = GlobalKey<NavigatorState>();

  PonygramRoutePath get currentConfiguration {
    if (show404) {
      return PonygramRoutePath.unknown();
    }
//    : PonygramRoutePath.details(books.indexOf(_selectedBook));
    if(page == 'account') {
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
 */
        if (show404 || page == 'home')
          MaterialPage(key: ValueKey('Home'), child: HomePage())
        else if (page == 'account')
          MaterialPage(key: ValueKey('Account'), child: AccountPage())
      ],
      onPopPage: (route, result) {
        if (!route.didPop(result)) {
          return false;
        }

        // Update the list of pages by setting _selectedBook to null
        show404 = false;
        page = '';
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
/*
  void _handleBookTapped(Book book) {
    _selectedBook = book;
    notifyListeners();
  }
 */
}
