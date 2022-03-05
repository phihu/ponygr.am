import 'package:flutter/material.dart';
import 'ponygram_route_path.dart';

class PonygramRouteInformationParser extends RouteInformationParser<PonygramRoutePath> {
  @override
  Future<PonygramRoutePath> parseRouteInformation(
      RouteInformation routeInformation) async {
    var test = Uri.tryParse(routeInformation.location ?? "");
    if(test == null){
      return PonygramRoutePath.home();
    }
    final uri = test;
    // Handle '/'
    if (uri.pathSegments.isEmpty) {
      return PonygramRoutePath.home();
    }
    if (uri.pathSegments.length == 1) {
      switch(uri.pathSegments[0]){
        case 'account':
          return PonygramRoutePath.account();
        case 'search':
          return PonygramRoutePath.search();
        case 'post':
          return PonygramRoutePath.post();
        case 'messages':
          return PonygramRoutePath.messages();
      }
      return PonygramRoutePath.home();
    }

    // Handle '/book/:id'
    if (uri.pathSegments.length == 2) {
      if (uri.pathSegments[0] != 'book') return PonygramRoutePath.unknown();
      var remaining = uri.pathSegments[1];
      var id = int.tryParse(remaining);
      if (id == null) return PonygramRoutePath.unknown();
//      return PonygramRoutePath.details(id);
      return PonygramRoutePath.home();
    }

    // Handle unknown routes
    return PonygramRoutePath.unknown();
  }

  @override
  RouteInformation restoreRouteInformation(PonygramRoutePath configuration) {
    if (configuration.isUnknown) {
      return const RouteInformation(location: '/404');
    }
    if (configuration.isAccountPage) {
      return const RouteInformation(location: '/account');
    }
    if (configuration.isMessagesPage) {
      return const RouteInformation(location: '/messages');
    }
    if (configuration.isSearchPage) {
      return const RouteInformation(location: '/search');
    }
    if (configuration.isPostPage) {
      return const RouteInformation(location: '/post');
    }
    if (configuration.isHomePage) {
      return const RouteInformation(location: '/');
    }
/*
    if (configuration.isDetailsPage) {
      return RouteInformation(location: '/post/${configuration.id}');
    }
*/
    return const RouteInformation(location: '/');
//    return null;
  }
}