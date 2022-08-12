Simple
---

This is the simplest sample using [Modus](https://modus-continens.com/).

## Usage

1. Analyse a Modusfile and checks the predicate kinds.
    ```shell
    modus check .
    ```

3. Print proof tree of a given query.
    ```shell
    modus proof . 'app'
    ```

4. Build images.
    ```shell
    modus build --json . 'app'
    ```

5. Create a tag TARGET_IMAGE that refers to SOURCE_IMAGE.
    ```shell
    docker tag sha256:xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx simple
    ```

6. Run the application.
    ```shell
    docker run -p 8080:8080 simple
    ```
