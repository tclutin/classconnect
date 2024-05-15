<?php

use Illuminate\Support\Facades\Route;
use App\Http\Controllers\AuthController;
use App\Http\Controllers\GroupController;
use App\Http\Controllers\ScheduleController;

Route::get('/', function () {
    return view('welcome');
});

Route::get('/signup', [AuthController::class, 'showSignupForm'])->name('signup');
Route::post('/signup', [AuthController::class, 'signup']);

Route::get('/login', [AuthController::class, 'showLoginForm'])->name('login');
Route::post('/login', [AuthController::class, 'login']);

// Выход пользователя
Route::post('/logout', [AuthController::class, 'logout'])->name('logout');

Route::get('/me', [AuthController::class, 'showUserInfo'])->name('me');

// Получение списка всех групп
Route::get('/groups', [GroupController::class, 'index'])->name('groups.index');

// Показывает форму для создания группы
Route::get('/groups/create', [GroupController::class, 'showCreateForm'])->name('groups.create_form');

// Создание группы
Route::post('/groups', [GroupController::class, 'create'])->name('groups.create');

Route::get('/schedules/upload', [ScheduleController::class, 'showForm'])->name('schedules.upload_form');
Route::post('/schedules', [ScheduleController::class, 'upload'])->name('schedules.upload');
