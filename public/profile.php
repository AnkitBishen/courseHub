<?php
include 'header.php';
?>


<main>
    <div class="profile-header">
        <div class="container">
            <div class="row align-items-center">
                <div class="col-md-2">
                    <img src="https://via.placeholder.com/150" alt="Profile Picture" class="img-fluid rounded-circle">
                </div>
                <div class="col-md-10">
                    <h1>John Doe</h1>
                    <p>Web Developer | Lifelong Learner</p>
                </div>
            </div>
        </div>
    </div>
    <div class="container">
        <div class="row">
            <div class="col-md-4">
                <h2>Personal Information</h2>
                <form>
                    <div class="mb-3">
                        <label for="name" class="form-label">Name</label>
                        <input type="text" class="form-control" id="name" value="John Doe">
                    </div>
                    <div class="mb-3">
                        <label for="email" class="form-label">Email</label>
                        <input type="email" class="form-control" id="email" value="john.doe@example.com">
                    </div>
                    <div class="mb-3">
                        <label for="bio" class="form-label">Bio</label>
                        <textarea class="form-control" id="bio" rows="3">Web developer passionate about creating user-friendly applications and continuously learning new technologies.</textarea>
                    </div>
                    <button type="submit" class="btn btn-primary">Update Profile</button>
                </form>
            </div>
            <div class="col-md-8">
                <h2>Enrolled Courses</h2>
                <div class="list-group">
                    <a href="#" class="list-group-item list-group-item-action">
                        <div class="d-flex w-100 justify-content-between">
                            <h5 class="mb-1">Web Development Fundamentals</h5>
                            <small>Progress: 60%</small>
                        </div>
                        <p class="mb-1">Learn the basics of HTML, CSS, and JavaScript to build modern websites.</p>
                        <small>8 weeks course</small>
                    </a>
                    <a href="#" class="list-group-item list-group-item-action">
                        <div class="d-flex w-100 justify-content-between">
                            <h5 class="mb-1">Introduction to Data Science</h5>
                            <small>Progress: 30%</small>
                        </div>
                        <p class="mb-1">Discover the world of data analysis, visualization, and machine learning.</p>
                        <small>10 weeks course</small>
                    </a>
                </div>
                <h2 class="mt-4">Achievements</h2>
                <div class="row">
                    <div class="col-md-4 mb-3">
                        <div class="card">
                            <img src="https://via.placeholder.com/100?text=HTML5" class="card-img-top" alt="HTML5 Badge">
                            <div class="card-body">
                                <h5 class="card-title">HTML5 Master</h5>
                                <p class="card-text">Completed all HTML5 modules with excellence.</p>
                            </div>
                        </div>
                    </div>
                    <div class="col-md-4 mb-3">
                        <div class="card">
                            <img src="https://via.placeholder.com/100?text=CSS3" class="card-img-top" alt="CSS3 Badge">
                            <div class="card-body">
                                <h5 class="card-title">CSS3 Wizard</h5>
                                <p class="card-text">Created stunning designs using advanced CSS3 techniques.</p>
                            </div>
                        </div>
                    </div>
                    <div class="col-md-4 mb-3">
                        <div class="card">
                            <img src="https://via.placeholder.com/100?text=JavaScript" class="card-img-top" alt="JavaScript Badge">
                            <div class="card-body">
                                <h5 class="card-title">JavaScript Ninja</h5>
                                <p class="card-text">Mastered core JavaScript concepts and DOM manipulation.</p>
                            </div>
                        </div>
                    </div>
                </div>
            </div>
        </div>
    </div>
</main>
