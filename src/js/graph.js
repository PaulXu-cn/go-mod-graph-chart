import * as d3 from 'd3';
import CSS from '../css/styles.css';

export default function graph() {
    let width = 1000, height = 1000;
    let colorScale, color, circles, lines, texts, arrowMarker, tickCount, simulation;
    let svg = d3.select('#app').append('svg').attr('width', width).attr('height', height);

    let nodes = [{ name: 'haha', value: 12 }, { name: '323', value: 32 }, { name: '我们', value: 24 }, { name: 'yes~', value: 43 }, { name: 'sukabulei', value: 18 }],
        links = [
            { 'source': 0, 'target': 3, weight: 1 },
            { 'source': 3, 'target': 4, weight: 1 },
            { 'source': 1, 'target': 2, weight: 0.5 },
        ];

    function dragstarted(d) {
        d3.select(this).raise().attr("stroke", "black");
        simulation.stop();
    }
    function dragged(d) {
        d3.select(this).attr("cx", d.x = d3.event.x).attr("cy", d.y = d3.event.y);
        ticked();
    }
    function dragended(d) {
        d3.select(this).attr("stroke", null);
        simulation.restart();
    }
    const drag = d3.drag()
        .on("start", dragstarted)
        .on("drag", dragged)
        .on("end", dragended);

    const render_init = function () {
        var defs = svg.append("defs");

        arrowMarker = defs.append("marker")
            .attr("id", "arrow")
            .attr("markerUnits", "strokeWidth")
            .attr("markerWidth", "12")
            .attr("markerHeight", "12")
            .attr("viewBox", "0 0 12 12")
            .attr("refX", "6")
            .attr("refY", "6")
            .attr("orient", "auto");

        var arrow_path = "M2,2 L10,6 L2,10 L6,6 L2,2";

        arrowMarker.append("path")
            .attr("d", arrow_path)
            .attr("fill", "black");

        lines = svg.selectAll('line').data(links)
            .enter().append('line')
            .attr('stroke', 'black')
            .attr('opacity', 0.8)
            // .attr('stroke-width', d => d.width * 5)
            .attr('stroke-width', 1)
            .attr("marker-end", "url(#arrow)");

        circles = svg.selectAll('circle').data(nodes)
            .enter().append('circle')
            .attr('r', 8)
            .attr('fill', function (d, i) { return color(i); })
            .call(drag);

        texts = d3.select('svg').selectAll('.circle-text').data(nodes)
            .enter().append('text').classed('circle-text', true).text(d => d.name)
            .attr('fill', 'gray')
            .attr('transform', `translate(10, 5)`);
    }

    tickCount = 1;

    function ticked() {
        lines.attr('x1', function (d) { return d.source.x; }).
            attr('y1', function (d) { return d.source.y; }).
            attr('x2', function (d) { if (0 < d.source.x - d.target.x) { return d.target.x + 8; } else { return d.target.x - 8; } }).
            attr('y2', function (d) { if (0 < d.source.y - d.target.y) { return d.target.y + 8; } else { return d.target.y - 8; } }).
            attr("marker-end", "url(#arrow)").
            attr('d', function (d) { return 'M ' + d.source.x + ' ' + d.source.y + ' L ' + d.target.x + ' ' + d.target.y }).
            classed('edgeline', true);

        circles.attr('cx', function (d) { return d.x; })
            .attr('cy', function (d) { return d.y; })

        texts.attr('x', function (d) { return d.x; }).attr('y', d => d.y);

        // if (tickCount++ > 1000) {
        //     return false;
        // }
    }

    function strengthFunc(link) {
        return link.weight * 0.01;
    }

    function distanceFunc(link) {
        return 200 * (1 / link.weight);
    }

    d3.json('/graph.json').then(data => {
        links = data.data.links;
        nodes = data.data.nodes;

        colorScale = d3.scaleDiverging(d3.interpolatePuOr).domain([0, (nodes.length - 1)]);
        color = d3.scaleDiverging(d3.interpolateRainbow).domain([0, nodes.length - 1])
        tickCount = 1;

        render_init();

        simulation = d3.forceSimulation(nodes)
            .force('charge', d3.forceManyBody)
            .force('center', d3.forceCenter(width / 2, height / 2))
            .force('link', d3.forceLink(links).strength(strengthFunc).distance(distanceFunc))
            .alphaTarget(0.3)
            .on('tick', ticked);
    })
}