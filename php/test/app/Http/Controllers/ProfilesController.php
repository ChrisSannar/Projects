<?php

namespace App\Http\Controllers;

use App\User;
// use Illuminate\Http\Request;

class ProfilesController extends Controller
{
    public function index($user) {
        // prints out whatever it is on the page
        // dd($user);

        // Find the user or fail gracefully
        $temp = User::findOrFail($user);
        // dd($temp);

        return view('profiles.index', [
            'user' => $temp,
        ]);
    }

    // in the case that we need to get back to the profile, lets get the current user
    public function redirect() {
        if (auth()->user()) {
            return redirect('/profile/' . auth()->user()->id);
        } else {
            return redirect('/login');
        }
    }
}
