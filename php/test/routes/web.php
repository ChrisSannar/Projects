<?php

use Illuminate\Support\Facades\Route;

/*
|--------------------------------------------------------------------------
| Web Routes
|--------------------------------------------------------------------------
|
| Here is where you can register web routes for your application. These
| routes are loaded by the RouteServiceProvider within a group which
| contains the "web" middleware group. Now create something great!
|
*/

Route::get('/', function () {
    return view('welcome');
});

Auth::routes();

Route::get('/profile', 'ProfilesController@redirect');
Route::get('/profile/{user}', 'ProfilesController@index')->name('profile.show');


Route::get('/home', function() {
    return redirect('/profile');
});

// Tasks
Route::get('/tasks', 'TaskController@index');
Route::get('/tasks/create', 'TaskController@create');
Route::post('/tasks', 'TaskController@store');

// Testing
Route::get('/test2', 'TestController@index');
Route::get('/test', 'TestController@api');
Route::put('/test', 'TestController@put');
