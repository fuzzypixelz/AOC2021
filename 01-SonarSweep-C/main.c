#include <stdint.h>
#include <stdlib.h>
#include <stdio.h>

size_t part_one(const char* filename) {
    char buffer[32] = { 0 };
    FILE* input = fopen(filename, "r");
    size_t increases = 0;
    size_t previous  = 0;
    size_t current   = 0;
    fscanf(input, "%zu", &previous);
    while (fscanf(input, "%zu", &current) != EOF) {
        if (current > previous) {
            increases++;
        }
        previous = current;
    }
    return increases;
}

int32_t scan_window(FILE* input, size_t* result) {
    size_t window[3] = { 0 };
    size_t backtrack = 0;
    for (size_t i = 0; i < 3; i++) {
        if (fscanf(input, "%zu", &window[i]) == EOF) {
            return EOF; // Stopped prematurely, no measurement.
        }
        if (i == 0) {
            // This is the position where we'd like to return
            // i.e. we only consumed one line.
            backtrack = ftell(input);
        }
    }
    *result = window[0] + window[1] + window[2];
    // We seek back two lines
    fseek(input, backtrack, SEEK_SET);
    return 0;
}

size_t part_two(const char* filename) {
    char buffer[32] = { 0 };
    FILE* input = fopen(filename, "r");
    size_t increases = 0;
    size_t previous  = 0;
    size_t current   = 0;
    // Assuming there is at least one window ...
    previous = scan_window(input, &previous);
    // TODO: refactor this?
    while (scan_window(input, &current) != EOF) {
        if (current > previous) {
            increases++;
        }
        previous = current;
    }
    return increases;
}

int main(void) {
    printf("%zu measurements are larger than the previous measurement!\n",
            part_one("input"));
    printf("%zu sums are larger than the previous sum!\n",
            part_two("input"));
    return EXIT_SUCCESS;
}
