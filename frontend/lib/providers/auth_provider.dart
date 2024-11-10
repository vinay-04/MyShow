import 'package:flutter_riverpod/flutter_riverpod.dart';
import 'package:myshow/models/user.dart';
import 'package:myshow/services/api_service.dart';
import 'package:shared_preferences/shared_preferences.dart';

final authProvider = StateNotifierProvider<AuthNotifier, User?>((ref) {
  return AuthNotifier(ref);
});

class AuthNotifier extends StateNotifier<User?> {
  final Ref _ref;
  AuthNotifier(this._ref) : super(null);

  Future<void> login(String username, String password) async {
    final user =
        await _ref.read(apiServiceProvider).loginUser(username, password);
    state = user;
  }

  Future<void> register(String username, String email, String password) async {
    final user = await _ref
        .read(apiServiceProvider)
        .registerUser(username, email, password);
    state = user;
  }

  Future<void> checkAuthStatus() async {
    final prefs = await SharedPreferences.getInstance();
    final token = prefs.getString('jwtToken');
    if (token != null) {
      try {
        await _ref.read(apiServiceProvider).getEvents();
        state = User(
            username: 'dummy',
            email: 'dummy'); // Dummy user to indicate logged in
      } catch (e) {
        state = null;
      }
    } else {
      state = null;
    }
  }

  void logout() async {
    final prefs = await SharedPreferences.getInstance();
    await prefs.remove('jwtToken');
    state = null;
  }
}

final apiServiceProvider = Provider<ApiService>((ref) {
  return ApiService();
});
