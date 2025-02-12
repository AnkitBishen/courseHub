<footer class="bg-light text-center text-lg-start mt-4">
    <div class="container p-4">
        <div class="row">
            <div class="col-lg-6 col-md-12 mb-4 mb-md-0">
                <h5 class="text-uppercase">Course Platform</h5>
                <p>
                    Empowering learners worldwide with high-quality online education.
                </p>
            </div>
            <div class="col-lg-3 col-md-6 mb-4 mb-md-0">
                <h5 class="text-uppercase">Links</h5>
                <ul class="list-unstyled mb-0">
                    <li><a href="index.html" class="text-dark">Home</a></li>
                    <li><a href="courses.html" class="text-dark">Courses</a></li>
                    <li><a href="about.html" class="text-dark">About</a></li>
                    <li><a href="#" class="text-dark">Contact</a></li>
                </ul>
            </div>
            <div class="col-lg-3 col-md-6 mb-4 mb-md-0">
                <h5 class="text-uppercase">Follow Us</h5>
                <ul class="list-unstyled mb-0">
                    <li><a href="#" class="text-dark">Facebook</a></li>
                    <li><a href="#" class="text-dark">Twitter</a></li>
                    <li><a href="#" class="text-dark">Instagram</a></li>
                    <li><a href="#" class="text-dark">LinkedIn</a></li>
                </ul>
            </div>
        </div>
    </div>
    <div class="text-center p-3" style="background-color: rgba(0, 0, 0, 0.2);">
        © 2023 Course Platform. All rights reserved.
    </div>
</footer>

<script src="https://cdn.jsdelivr.net/npm/bootstrap@5.3.0-alpha1/dist/js/bootstrap.bundle.min.js"></script>
<script src="https://ajax.googleapis.com/ajax/libs/angularjs/1.8.2/angular.min.js"></script>
<script src="assets/js/app.js"></script>
<script src="assets/js/headerScript.js"></script>
<?php
    if (isset($js)) {
        foreach ($js as $file) {
            echo '<script src="' . $file . '"></script>';
        }
    }
?>
</body>
</html>