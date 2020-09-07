<?php

namespace App\Http\Controllers;

use Illuminate\Http\Request;

class TaskController extends Controller
{
    public function __construct() {
        $this->middleware('auth');  // sets the 'auth' middleware of all other functions here

        // Now if we try to access any of these, it'll just redirect to the signin page
    }

    // The Tasks main page
    public function index() {

        // dd(auth()->user()->tasks);
        // Display the authorized users tasks
        return view('tasks.index', [
            'tasks' => auth()->user()->tasks,
        ]);
    }

    // Create a new Task
    public function create() {
        return view('tasks.create');
    }

    // A request to save a new task
    public function store() {

        // Validate the data from the request
        $data = request()->validate([
            'text' => 'required'
        ]);
        $data['completed'] = false;

        auth()->user()->tasks()->create($data); // Grabs the authenticated user and creates the task on their account

        // \App\Task::create($data);

        // \App\Task::create([
        //     'text' => $data['text'],
        //     'user_id' => 1,
        //     'completed' => false,
        // ]);

        // $task = new \App\Task();
        // $task->text = $data['text'];
        // $task->user_id = 1;
        // $task->completed = FALSE;
        // $task->save();

        // dd(request()->all());

        return redirect('/tasks');
    }
}
