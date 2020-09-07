<?php

namespace App\Http\Controllers;

use App\User;
// use Illuminate\Http\Request;

/*
    This file is how we get the data and render it to the page.
    Mainly each Controller is used to define some kind of interaction/object then how that object works
    Ex. How do we get the profile form the database?
        What data is being rendered to the page?
        How do we organize that data?
        What view do we 
*/

class ProfilesController extends Controller
{
    public function index($user) {
        // prints out whatever it is on the page
        // dd($user);

        // Find the user or fail gracefully
        $temp = User::findOrFail($user);
        // dd($temp);

        return view('home', [
            'user' => $temp,
        ]);
    }
}
