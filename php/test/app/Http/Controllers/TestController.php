<?php

namespace App\Http\Controllers;

use App\Task;
use App\Http\Resources\TestResource;
use Illuminate\Http\Request;


/*
    This file is how we get the data and render it to the page.
    Mainly each Controller is used to define some kind of interaction/object then how that object works
    Ex. How do we get the profile form the database?
        What data is being rendered to the page?
        How do we organize that data?
        What view do we 
*/

class TestController extends Controller
{
    public function index() {
        return view('test');
    }

    public function api() {
        return TestResource::collection(Task::all());
    }
}
