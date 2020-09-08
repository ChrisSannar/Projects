<?php

namespace App\Http\Controllers;

use App\Task;
use App\Http\Resources\TestResource;
use Illuminate\Http\Request;
use Validator;


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

        // Use a resource to format the tasks before sending them out.
        return TestResource::collection(Task::all());
    }

    public function put() {

        // Wow this took too long. Learned a ton from it thought.
        
        // Parse the data from JSON into php format
        $data = json_decode(request()->getContent(), true);
        
        // Since we have an array of items, we need to update each one
        // I understand this can be expensive, but at this point it's working
        for($i = 0; $i < count($data); $i++) {

            // Grab the task that has both the user id and the Task id
            $task = Task::where('user_id', auth()->user()->id)->where('id', $data[$i]['id']);
            
            // Update that specific task with the matching properties
            $task->update([
                'text' => $data[$i]['text'],
                'completed' => $data[$i]['completed'],
            ]);
        }

        return "OK";
            
        // As you can see below, it took me a while to figure this one out...


        // return json_encode($tasks);
        
        // dd($data);
        // $tasks = Task::where('user_id', auth()->user()->id)->update($data);
        // dd($data);
        // Validate the incoming request
        // $rules = [
        //     'text' => 'required',
        //     'completed' => 'required',
        // ];
        // $validator = Validator::make($data, $rules);

        // if ($validator->passes()) {
        //     $tasks = auth()->user()->tasks();
        //     return $tasks;
        // } else {
        //     error_log('mark');
        //     return json_encode($validator->errors()->all());
        // }
        // $currenTasks = auth()->user()->tasks();
        
        // $currenTasks->update($data);
        // $currenTasks->push();

        // $tasks = $data;
        // foreach($data as $task) {
        //     var_dump($task);
        // }
        // dd($data);

        // dd($currenTasks);
        // return $currenTasks;

    }
}
