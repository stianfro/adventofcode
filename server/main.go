package main

import (
	"context"
	"log"
	"net"

	aoc2015day1 "github.com/stianfro/adventofcode/2015/day-1"
	aoc2015day2 "github.com/stianfro/adventofcode/2015/day-2"
	proto2015 "github.com/stianfro/adventofcode/gen/go/proto/v2015"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type server struct {
	proto2015.SolutionServiceServer
}

func (s *server) Day1(ctx context.Context, in *proto2015.SolutionRequest) (*aoc2015.SolutionResponse, error) {
	input := in.Input

	return &proto2015.SolutionResponse{
		AnswerPartOne: aoc2015day1.PartOne(input),
		AnswerPartTwo: aoc2015day1.PartTwo(input),
	}, nil
}

func (s *server) Day2(ctx context.Context, in *proto2015.SolutionRequest) (*aoc2015.SolutionResponse, error) {
	input := in.Input

	return &proto2015.SolutionResponse{
		AnswerPartOne: aoc2015day2.PartOne(input),
		AnswerPartTwo: aoc2015day2.PartTwo(input),
	}, nil
}

func main() {
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal("failed to serve:", err)
	}

	s := grpc.NewServer()
	reflection.Register(s)

	proto2015.RegisterSolutionServiceServer(s, &server{})
	if err := s.Serve(listener); err != nil {
		log.Fatal("failed to serve:", err)
	}
}
